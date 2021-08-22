package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/namnguyen191/themuzix-golang-rest-api/service"
)

type SongController interface {
	All(*gin.Context)
	FindByID(*gin.Context)
	Insert(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type songController struct {
	songService service.SongService
	jwtService  service.JWTService
}

func (c *songController) All(_ *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *songController) FindByID(_ *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *songController) Insert(_ *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *songController) Update(_ *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *songController) Delete(_ *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func NewSongController(songSer service.SongService, jwtSer service.JWTService) SongController {
	return &songController{
		songService: songSer,
		jwtService:  jwtSer,
	}
}
