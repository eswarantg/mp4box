package mp4box

//HandlerBox -
/*
aligned(8) class HandlerBox extends FullBox(‘hdlr’, version = 0, 0) { unsigned int(32) pre_defined = 0;
unsigned int(32) handler_type;
const unsigned int(32)[3] reserved = 0;
   string   name;
}
*/
type HandlerBox struct {
	FullBox
}
