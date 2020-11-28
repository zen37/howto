Option Explicit
 
Private WithEvents oItems As Outlook.Items
 
 
Private Sub Application_Startup()
 
Dim oApp As Outlook.Application
Dim oNS As Outlook.NameSpace
 
Set oApp = Outlook.Application
Set oNS = oApp.GetNamespace("MAPI")
 
Set oItems = oNS.GetDefaultFolder(olFolderInbox).Items
 
Debug.Print "Application_Startup started " & Now()
 
End Sub
 
Private Sub oItems_ItemAdd(ByVal item As Object)
 
Dim oMail As Outlook.MailItem
 
If TypeName(item) = "MailItem" Then
 
 
    Set oMail = item
 
    Debug.Print "New e-mail received " & Now()
 
'Send_email
    Debug.Print "oMail.ConversationID = " & oMail.ConversationID
   
    Debug.Print "oMail.ConversationIndex = " & oMail.ConversationIndex
   
    Debug.Print "oMail.ConversationTopic = " & oMail.ConversationTopic
   
    
    Debug.Print oMail.Subject
    Debug.Print oMail.Body
   
    
    'Debug.Print "Length: " & Len(oMail.Body)
 
End If
 
End Sub