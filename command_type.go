package mySQL

type CommandType string

const (
	Get    CommandType = "get"
	List   CommandType = "list"
	Insert CommandType = "insert"
	Update CommandType = "update"
	Delete CommandType = "delete"
)

type CommandTypeBase struct {
	commandType CommandType
}

func (d *CommandTypeBase) SetCommandType(t CommandType) {
	d.commandType = t
}

func (d *CommandTypeBase) IsGet() bool {
	return d.commandType == Get
}

func (d *CommandTypeBase) IsList() bool {
	return d.commandType == List
}

func (d *CommandTypeBase) IsInsert() bool {
	return d.commandType == Insert
}

func (d *CommandTypeBase) IsUpdate() bool {
	return d.commandType == Update
}

func (d *CommandTypeBase) IsDelete() bool {
	return d.commandType == Delete
}
