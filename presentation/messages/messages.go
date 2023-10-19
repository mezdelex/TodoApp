package messages

import "fmt"

type Messages struct{}
type Status struct{}

// Messages
func (_ Messages) CollectionEmptyMessage(itemName string) string { return fmt.Sprintf("The %s collection is empty.", itemName) }
func (_ Messages) IdConflictErrorMessage(itemName string) string { return fmt.Sprintf("The route id and the %s's id are not equal.", itemName) }
func (_ Messages) ItemCreatedSuccessfullyMessage(item interface{}) string { return fmt.Sprintf("%v was created successfully.", item) }
func (_ Messages) ItemDeletedSuccessfullyMessage(itemName string, id uint) string { return fmt.Sprintf("%s with id %d was deleted successfully.", itemName , id) }
func (_ Messages) LoggedInSuccessfullyMessage() string { return fmt.Sprint("Logged in successfully.") }
func (_ Messages) ParsingErrorMessage(itemName string) string { return fmt.Sprintf("The provided %s could not be parsed.", itemName) }
func (_ Messages) ReturningItemsSuccessfullyMessage(length int, itemName string) string { return fmt.Sprintf("Returning %d %s(s).", length, itemName) }
func (_ Messages) RouteFormatErrorMessage(parameter string) string { return fmt.Sprintf("Incorrect format in route's '%s' parameter.", parameter) }
func (_ Messages) ValuesUpdatedSuccessfullyMessage(itemName string) string { return fmt.Sprintf("%s values updated successfully.", itemName) }

// Status
func (_ Status) Error() string   { return "error" }
func (_ Status) Success() string { return "success" }
