package employees

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	if c.Request().Method != echo.GET {
		return c.String(405, "StatusMethodNotAllowed")
	}

	emps, err := AllEmps()
	if err != nil {
		return c.String(500, "StatusInternalServerError")
	}
	return c.Render(http.StatusOK, "employees.gohtml", emps)
}

func Show(c echo.Context) error {
	if c.Request().Method != echo.GET {
		return c.String(405, "StatusMethodNotAllowed")
	}

	emp, err := OneEmp(c)
	if err != nil {
		return c.String(500, "StatusInternalServerError")
	}
	return c.Render(http.StatusOK, "show.gohtml", emp)
}

func Create(c echo.Context) error {
	return c.Render(http.StatusOK, "create.gohtml", nil)
}

func CreateProcess(c echo.Context) error {
	if c.Request().Method != echo.POST {
		return c.String(405, "StatusMethodNotAllowed")
	}

	emp, err := PutEmp(c)
	if err != nil {
		return c.String(500, "StatusInternalServerError")
	}

	return c.Render(http.StatusOK, "created.gohtml", emp)
}

func Update(c echo.Context) error {
	if c.Request().Method != echo.GET {
		return c.String(405, "StatusMethodNotAllowed")
	}

	emp, err := OneEmp(c)
	if err != nil {
		return c.String(500, "StatusInternalServerError")
	}

	return c.Render(http.StatusOK, "update.gohtml", emp)
}

func UpdateProcess(c echo.Context) error {
	if c.Request().Method != echo.POST {
		return c.String(405, "StatusMethodNotAllowed")
	}

	emp, err := UpdateEmp(c)
	if err != nil {
		return c.String(500, "StatusInternalServerError")
	}

	return c.Render(http.StatusOK, "updated.gohtml", emp)
}

func DeleteProcess(c echo.Context) error {
	if c.Request().Method != echo.GET {
		return c.String(405, "StatusMethodNotAllowed")
	}

	err := DeleteEmp(c)
	if err != nil {
		return c.String(400, "StatusBadRequest")
	}
	return c.Redirect(303, "/emps")
}
