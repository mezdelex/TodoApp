package messages

import "fmt"

type Messages struct{}
type Status struct{}

// Messages
func (m Messages) CollectionEmptyMessage(itemName string) string { return fmt.Sprintf("The %s collection is empty.", itemName) }
func (m Messages) ItemCreatedMessage(item interface{}) string { return fmt.Sprintf("%v was created successfully.", item) }
func (m Messages) ParsingErrorMessage(itemName string) string { return fmt.Sprintf("The provided %s could not be parsed.", itemName) }
func (m Messages) ReturningItemsMessage(length int, itemName string) string { return fmt.Sprintf("Returning %d %s(s).", length, itemName) }
func (m Messages) RouteFormatErrorMessage(parameter string) string { return fmt.Sprintf("Incorrect format in route's '%s' parameter.", parameter) }
func (m Messages) UpdateIdsConflictMessage(itemName string) string { return fmt.Sprintf("The route id and the %s's id are not equal.", itemName) }
func (m Messages) ValuesUpdatedSuccessfullyMessage(itemName string) string { return fmt.Sprintf("%s values updated successfully.", itemName) }
func (m Messages) ItemDeletedSuccessfullyMessage(itemName string, id uint) string { return fmt.Sprintf("%s with id %d was deleted successfully.", itemName , id) }

// Status
func (s Status) Error() string   { return "error" }
func (s Status) Success() string { return "success" }
