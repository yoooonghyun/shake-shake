package app

import (
	"net/http"
	domain "shake-shake/src/domain/member"
	"shake-shake/src/service"

	"github.com/gin-gonic/gin"
)

type MemberApp struct {
	memberSvc *service.MemberService
}

func CreateMemberApp(router *gin.Engine) (*MemberApp, error) {
	memberSvc, err := service.CreateMemberService()
	app := MemberApp{memberSvc: memberSvc}

	if err != nil {
		return nil, err
	}

	router.POST("/member", app.createMember)

	router.GET("/members", app.getMembers)

	router.GET("/members/:memberId", app.getMember)

	router.POST("/shake-shake", app.shakeShake)

	return &app, nil
}

// @Summary Create member.
// @Description Create new member.
// @Success 200
// @Failure 400
// @Failure 500
// @Router /memeber [post]
func (app *MemberApp) createMember(c *gin.Context) {
	body := new(domain.Member)
	err := c.Bind(body)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err = app.memberSvc.Create(c, body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

// @Summary Get members.
// @Description Get members array.
// @Success 200 {array} domain.Member
// @Failure 500
// @Router /members [get]
func (app *MemberApp) getMembers(c *gin.Context) {
	members, err := app.memberSvc.ReadMany(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, members)
}

// @Summary Get a member.
// @Description Get a member w/ member id.
// @Success 200 {object} domain.Member
// @Failure 500
// @Router /members/{memberId} [get]
func (app *MemberApp) getMember(c *gin.Context) {
	memberId := c.Param("memberId")
	member, err := app.memberSvc.ReadOne(c, memberId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, member)
}

// @Summary shake-shake.
// @Description Mix members' group and return.
// @Router /shake-shake [post]
func (app *MemberApp) shakeShake(c *gin.Context) {
}
