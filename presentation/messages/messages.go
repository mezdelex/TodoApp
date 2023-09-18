package messages

import "fmt"

type Messages struct{}
type Status struct{}

// Messages
func (_ Messages) CollectionEmptyMessage(itemName string) string { return fmt.Sprintf("The %s collection is empty.", itemName) }
func (_ Messages) ItemCreatedMessage(item interface{}) string { return fmt.Sprintf("%v was created successfully.", item) }
func (_ Messages) ParsingErrorMessage(itemName string) string { return fmt.Sprintf("The provided %s could not be parsed.", itemName) }
func (_ Messages) ReturningItemsMessage(length int, itemName string) string { return fmt.Sprintf("Returning %d %s(s).", length, itemName) }
func (_ Messages) RouteFormatErrorMessage(parameter string) string { return fmt.Sprintf("Incorrect format in route's '%s' parameter.", parameter) }
func (_ Messages) IdConflictMessage(itemName string) string { return fmt.Sprintf("The route id and the %s's id are not equal.", itemName) }
func (_ Messages) ValuesUpdatedSuccessfullyMessage(itemName string) string { return fmt.Sprintf("%s values updated successfully.", itemName) }
func (_ Messages) ItemDeletedSuccessfullyMessage(itemName string, id uint) string { return fmt.Sprintf("%s with id %d was deleted successfully.", itemName , id) }

// Status
func (_ Status) Error() string   { return "error" }
func (_ Status) Success() string { return "success" }
