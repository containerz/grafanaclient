package grafanaclient

import "github.com/stretchr/testify/assert"
import "testing"

var url = "http://localhost:3000"

var ds = DataSource{Name: "testme",
	Type:      "influxdb_08",
	Access:    "direct",
	Url:       "http://localhost:8086",
	User:      "root",
	Password:  "root",
	Database:  "test",
	IsDefault: true}

var dashboard = `{
        "id": null,
        "title": "Production Overview",
        "tags": [ "templated" ],
        "timezone": "browser",
        "rows": [
          {
          }
        ],
        "schemaVersion": 6,
        "version": 0
      }`

// func Test_BadConnect(t *testing.T) {
// 	testDB := NewInfluxDB()
// 	err := testDB.InitSession("locallhost", "8087", "testdb", "admin", "admin")
// 	assert.NotNil(t, err, "We are expecting error and didn't got one")
// }

func Test_DoLogon(t *testing.T) {
	session := NewSession("admin", "admin", url)
	err := session.DoLogon()
	assert.Nil(t, err, "We are expecting no error and got one")
}

func Test_CreateDataSource(t *testing.T) {
	session := NewSession("admin", "admin", url)
	err := session.DoLogon()
	assert.Nil(t, err, "We are expecting no error and got one")
	err = session.CreateDataSource(ds)
	assert.Nil(t, err, "We are expecting no error and got one when creating DataSource")
}

func Test_GetDataSourceList(t *testing.T) {
	session := NewSession("admin", "admin", url)
	err := session.DoLogon()
	assert.Nil(t, err, "We are expecting no error and got one when Login")
	dslist, err := session.GetDataSourceList()
	assert.Nil(t, err, "We are expecting no error and got one getting DataSource")
	var check bool
	for _, ds := range dslist {
		if ds.Name == "testme" {
			check = true
		}
	}

	assert.Equal(t, true, check, "We didn't find the test datasource")
}

func Test_GetDataSource(t *testing.T) {
	session := NewSession("admin", "admin", url)
	err := session.DoLogon()
	assert.Nil(t, err, "We are expecting no error and got one when Login")

	resDs, err := session.GetDataSource("testme")

	assert.Equal(t, "testme", resDs.Name, "We are expecting to retrieve testme DataSource and didn't get it")
}

func Test_GetDashboard(t *testing.T) {
	session := NewSession("admin", "admin", url)
	err := session.DoLogon()
	assert.Nil(t, err, "We are expecting no error and got one when Login")

	dashboard := session.GetDashboard("new-dashboard")
	assert.NotNil(t, dashboard, "We are expecting to receive a dashboard")

}

func Test_CreateDashboard(t *testing.T) {
	session := NewSession("admin", "admin", url)
	err := session.DoLogon()
	assert.Nil(t, err, "We are expecting no error and got one when Login")

	session.UploadDashboard(dashboard)
}

// func Test_DeleteDataSource(t *testing.T) {
// 	session := NewSession("admin", "admin", url)
// 	err := session.DoLogon()
// 	assert.Nil(t, err, "We are expecting no error and got one when Login")

// 	resDs, err := session.GetDataSource("testme")

// 	err = session.DeleteDataSource(resDs)
// 	assert.Nil(t, err, "We are expecting no error and got one when Deleting")
// }
