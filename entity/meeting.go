package entity


import (
	"encoding/json"
	"fmt"
	"time"
	"os"
	"strings"
	"errors"
)

type Meeting struct {
	// The title of meeting, used to identify a meeting
	Title string

	Sponsor string

	// Participators of meeting
	Participators []string
	
	// Start time
	StartTime time.Time
	
	// End time
	EndTime time.Time
}

func (m *Meeting) hasParticipator(user string) bool {
	for _, participator := range m.Participators {
		if user == participator {
			return true
		}
	}
	return false
}

var meetingList []Meeting
var meetingFileName = "meetings.txt"
var timeFormat = "2006-01-02T15:04:05"

func loadMeetingsFromFile(filename string) error {
	file, err := os.Open(meetingFileName)
	if err != nil {
		return nil
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	if err := dec.Decode(&meetingList); err != nil {
		return err
	}
	return nil
}

func saveMeetingsToFile(filename string) error {
	file, err := os.OpenFile(meetingFileName, os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	if err := enc.Encode(meetingList); err != nil {
		return err
	}
	return nil
}

// get the meeting list that the user participates or sponsors
func getMyMeetings(meetings []Meeting, user string) []Meeting {
	var result []Meeting
	for _, meeting := range meetings {
		if meeting.hasParticipator(user) || meeting.Sponsor == user {
			result = append(result, meeting)
		}
	}
	return result
}

func getMySponsoredMeetings(meetings []Meeting, user string) []Meeting {
	var result []Meeting
	for _, meeting := range meetings {
		if meeting.Sponsor == user {
			result = append(result, meeting)
		}
	}
	return result
}


func getMyParticipatedMeetings(meetings []Meeting, user string) []Meeting {
	var result []Meeting
	for _, meeting := range meetings {
		if meeting.hasParticipator(user) {
			result = append(result, meeting)
		}
	}
	return result
}

func getMeetingByTitle(meetings []Meeting, title string) (Meeting, int) {
	for index, meeting := range meetings {
		if meeting.Title == title {
			return meeting, index
		}
	}
	return Meeting{}, -1
}

// verified whether all participators have been registered
func verifyParticipators(users UserArray, participators []string) error {
	var found bool
	for _, part := range participators {
		found = false
		for _, user := range users {
			if user.Username == part {
				found = true
			}
		}
		if !found {
			return errors.New("participator " + part + " has not been registered")
		}
	}
	return nil
}

// check if there are any meeting conflicting with the period from st to et
func checkConflict(meetings []Meeting, st time.Time, et time.Time) error {
	for _, meeting := range meetings {
		if !(meeting.StartTime.After(et) || meeting.StartTime.Equal(et) || meeting.EndTime.Before(st) || meeting.EndTime.Equal(st) ) {
			return errors.New(fmt.Sprintf("%v to %v conflicts with meeting %s", st, et, meeting.Title))
		}
	}
	return nil
}

// check whether there are any dulplicated participator
func checkDulplicate(participators []string) error {
	states := make(map[string]bool)
	for _, user := range participators {
		if _, exist := states[user]; exist {
			return errors.New("dulplicated participator " + user)
		}
		states[user] = true
	}
	return nil
}

// check if all participators have participated the meeting and return the list of participators after removing those participators
func checkParticipate(curMeeting Meeting, participators []string) (error, []string) {
	states := make(map[string]bool)
	for _, user := range curMeeting.Participators {
		states[user] = true
	}
	for _, user := range participators {
		if user == curMeeting.Sponsor {
			return errors.New("you can not remove yourself from the meeting or use cancel or quit command instead"), []string{}
		}
		if _, exist := states[user]; !exist {
			return errors.New("the participator " + user + " you want to remove does not participate the meeting"), []string{}
		}
		states[user] = false
	}
	var removedList []string 
	for _, user := range curMeeting.Participators {
		if states[user] {
			removedList =  append(removedList, user)
		}
	} 
	return nil, removedList

}

func CreateMeeting(title string, participatorsStr string, stStr string, etStr string) error {
	// parse parameters
	participators := strings.Split(participatorsStr, ",")
	st, err := time.Parse(timeFormat, stStr)
	if err != nil {
		return err
	}
	et, err := time.Parse(timeFormat, etStr)
	if err != nil {
		return err
	}

	// check if there is any participators
	if len(participators) == 0 {
		return errors.New("participators list is empty, you should specify at least one participator")
	}

	// check time
	if !et.After(st) {
		return errors.New("invalid time, end time is before start time")
	}

	// check logined and get user lists
	users, err := QueryUser()
	if err != nil {
		return err
	}
	curUser, err := GetCurUser()
	if err != nil {
		return err
	}
	sponsor := curUser.Username

	// load meetings list
	err = loadMeetingsFromFile(meetingFileName)
	if err != nil {
		return err
	}

	// verify whether all participators have been registered
	err = verifyParticipators(users, append(participators, sponsor))
	if err != nil {
		return err
	}
	// check dulplicated
	err = checkDulplicate(append(participators, sponsor))
	if err != nil {
		return err
	}

	// check title
	meetings := getMySponsoredMeetings(meetingList, sponsor)
	for _, meeting := range meetings {
		if meeting.Title == title {
			return errors.New("the title " + title + " already exists")
		}
	}

	// check conflicts
	for _, user := range append(participators, sponsor) {
		meetings = getMyMeetings(meetingList, user)
		err = checkConflict(meetings, st, et)
		if err != nil {
			return err
		}
	}

	// create new meeting
	newMeeting := Meeting{title, sponsor, participators, st, et}
	meetingList = append(meetingList, newMeeting)

	// save changes to file
	if err = saveMeetingsToFile(meetingFileName); err != nil {
		return err
	}

	return nil
}


// add participarors to the meeting
func AddParticipators(title string, participatorsStr string) error {
	// parse parameters
	participators := strings.Split(participatorsStr, ",")

	// check logined get user list
	users, err := QueryUser()
	if err != nil {
		return err
	}
	curUser, err := GetCurUser()
	if err != nil {
		return err
	}
	sponsor := curUser.Username

	// load meetings list
	err = loadMeetingsFromFile(meetingFileName)
	if err != nil {
		return err
	}

	// get the meeting 
	meetings := getMySponsoredMeetings(meetingList, sponsor)
	curMeeting, index := getMeetingByTitle(meetings, title)
	if index == -1 {
		return errors.New("you do not sponsor the meeting or the meeting does not exist")
	}

	// verify whether all participators have been registered
	err = verifyParticipators(users, participators)
	if err != nil {
		return err
	}

	// check dulplicated
	err = checkDulplicate(append(append(curMeeting.Participators, sponsor), participators...))
	if err != nil {
		return err
	}

	// check conflicts
	for _, user := range participators {
		meetings = getMyMeetings(meetingList, user)
		err = checkConflict(meetings, curMeeting.StartTime, curMeeting.EndTime)
		if err != nil {
			return err
		}
	}

	// add participators to current meeting
	curMeeting.Participators = append(curMeeting.Participators, participators...)

	// save changes to file
	if err = saveMeetingsToFile(meetingFileName); err != nil {
		return err
	}

	return nil
}


// need user.Equal(user)
func RemoveParticipators(title string, participatorsStr string) error {
	// parse parameters
	participators := strings.Split(participatorsStr, ",")

	// check logined get user list
	users, err := QueryUser()
	if err != nil {
		return err
	}
	curUser, err := GetCurUser()
	if err != nil {
		return err
	}
	sponsor := curUser.Username

	// load meetings list
	err = loadMeetingsFromFile(meetingFileName)
	if err != nil {
		return err
	}

	// get the meeting 
	meetings := getMySponsoredMeetings(meetingList, sponsor)
	curMeeting, index := getMeetingByTitle(meetings, title)
	if index == -1 {
		return errors.New("you do not sponsor the meeting or the title is not exist")
	}

	// verify whether all participators have been registered
	err = verifyParticipators(users, participators)
	if err != nil {
		return err
	}

	// check dulplicated
	err = checkDulplicate(participators)
	if err != nil {
		return err
	}

	// check whether all participators can be removed from current meeting
	err, removedList := checkParticipate(curMeeting, participators)
	if err != nil {
		return err
	}

	// remove participators from current meeting
	if len(removedList) == 0 {
		meetingList = append(meetingList[:index], meetingList[index + 1:]...)
	} else {
		curMeeting.Participators = removedList
	}

	// save changes to file
	if err = saveMeetingsToFile(meetingFileName); err != nil {
		return err
	}

	return nil
}

//query meetings of the sponsor by time
func QueryMeetingsByTime(stStr string, etStr string) (error, []Meeting) {
	// parse parameters
	st, err := time.Parse(timeFormat, stStr)
	if err != nil {
		return err, []Meeting{}
	}
	et, err := time.Parse(timeFormat, etStr)
	if err != nil {
		return err, []Meeting{}
	}

	// check time
	if !et.After(st) {
		return errors.New("invalid time, end time is before start time"), []Meeting{}
	}

	// check logined get user list
	_, err = QueryUser()
	if err != nil {
		return err, []Meeting{}
	}
	curUser, err := GetCurUser()
	if err != nil {
		return err, []Meeting{}
	}
	sponsor := curUser.Username

	// load meetings list
	err = loadMeetingsFromFile(meetingFileName)
	if err != nil {
		return err, []Meeting{}
	}

	// execute query
	meetings := getMyMeetings(meetingList, sponsor)
	var result []Meeting
	for _, meeting := range meetings {
		if (meeting.StartTime.After(et) || meeting.StartTime.Equal(et)) || (meeting.EndTime.Before(st) || meeting.EndTime.Equal(st)) {
			continue
		} else {
			result = append(result, meeting)
		}
	}

	return nil, result
}

func CancelMeeting(title string) error {
	// check logined get user list
	_, err := QueryUser()
	if err != nil {
		return err
	}
	curUser, err := GetCurUser()
	if err != nil {
		return err
	}
	sponsor := curUser.Username

	// load meetings list
	err = loadMeetingsFromFile(meetingFileName)
	if err != nil {
		return err
	}

	// get the meeting 
	meetings := getMySponsoredMeetings(meetingList, sponsor)
	_, index := getMeetingByTitle(meetings, title)
	if index == -1 {
		return errors.New("you do not sponsor the meeting or the meeting does not exist")
	}

	// cancel the meeting
	meetingList = append(meetingList[:index], meetingList[index + 1:]...)

	// save changes to file
	if err = saveMeetingsToFile(meetingFileName); err != nil {
		return err
	}
	return nil
}

func QuitMeeting(title string) error {
	// check logined get user list
	_, err := QueryUser()
	if err != nil {
		return err
	}
	curUser, err := GetCurUser()
	if err != nil {
		return err
	}
	sponsor := curUser.Username

	// load meetings list
	err = loadMeetingsFromFile(meetingFileName)
	if err != nil {
		return err
	}

	// get the meeting 
	meetings := getMyParticipatedMeetings(meetingList, sponsor)
	curMeeting, index := getMeetingByTitle(meetings, title)
	if index == -1 {
		return errors.New("you do not sponsor the meeting or the meeting does not exist")
	}

	// quit the meeting
	myIndex := 0
	for id, user := range curMeeting.Participators {
		if user == sponsor {
			myIndex = id
			break
		}
	}
	curMeeting.Participators = append(curMeeting.Participators[:myIndex], curMeeting.Participators[myIndex + 1:]...)
	if len(curMeeting.Participators) == 0 {
		meetingList = append(meetingList[:index], meetingList[index + 1:]...)
	}

	// save changes to file
	if err = saveMeetingsToFile(meetingFileName); err != nil {
		return err
	}
	return nil
}

func ClearMeetings() error {
	// check logined get user list
	_, err := QueryUser()
	if err != nil {
		return err
	}
	curUser, err := GetCurUser()
	if err != nil {
		return err
	}
	sponsor := curUser.Username

	// load meetings list
	err = loadMeetingsFromFile(meetingFileName)
	if err != nil {
		return err
	}

	// clear all sponsored meetings
	var newMeetingList []Meeting
	states := make(map[string]bool)
	for _, meeting := range meetingList {
		if meeting.Sponsor != sponsor {
			states[meeting.Title] = true
		} else {
			states[meeting.Title] = false
		}
	}
	for _, meeting := range meetingList {
		if states[meeting.Title] {
			newMeetingList = append(newMeetingList, meeting)
		}
	}
	meetingList = newMeetingList

	// save changes to file
	if err = saveMeetingsToFile(meetingFileName); err != nil {
		return err
	}
	return nil
}
