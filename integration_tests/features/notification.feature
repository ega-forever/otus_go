# file: features/notification.feature

# http://localhost:8088/
# http://registration_service:8088/

Feature: Email notification sending
	As API client of registration service
	In order to understand that the user was informed about registration
	I want to receive event from notifications queue

	Scenario: Notification event is received
		When I create new event "super event" and request to "localhost:8082"
		Then the record should be created
		And I receive event with text "super event"
