package cube

const (
	CUBE_ORDER          = 5
	MAGIC_NUMBER        = (CUBE_ORDER * (SEQUENCE_SIZE + 1)) / 2       // 315
	MAGIC_NUMBER_AMOUNT = 3*(CUBE_ORDER*CUBE_ORDER) + 6*CUBE_ORDER + 4 // 109
	SEQUENCE_SIZE       = CUBE_ORDER * CUBE_ORDER * CUBE_ORDER         // 125
)
