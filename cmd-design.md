
Usage:
	Agenda [Command] [Flags]

---
Command cm:
> 	cm 	create a meeting, you should specify the title, start time, end time and participators of the 
 		meeting. e.g., cm -t newmeeting -st "2018-10-10 09:10:00" -et "2018-10-10 09:50:00" -p
 		user1,user2,user3 

Flags:
	-t  	--title 		string	title of meeeting
	-s 	--starttime		string 	start time of meeting, format required: "2018-10-01 08:00:00"
	-e 	--endtime   	string  end time of meeting, format required: "2018-10-01 08:00:00"
	-p		--partipators 	string	a list of partipators (identify each user by his name), use comma to separate

---
Command mm:
> 	mm 	add/remove participators from the meeting, you should specify the title of the meeting and a list of 
 		parptipators that you want to add/remove, use -r to remove those participators from the meeting.
 		e.g., mm mymeeting user1,user2 (add user1 and user2 to mymeeting if they are not in the meeting) 

Flags:		
	-t  	--title 		string	title of meeeting
	-p	--partipators 		string	a list of partipators (identify each user by his name), use comma to separate							
	-r 	--remove		bool	remove participator(s)

