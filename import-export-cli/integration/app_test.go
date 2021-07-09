/*
*  Copyright (c) WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 Inc. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package integration

import (
	"testing"

	"github.com/wso2/product-apim-tooling/import-export-cli/integration/apim"
	"github.com/wso2/product-apim-tooling/import-export-cli/integration/base"
	"github.com/wso2/product-apim-tooling/import-export-cli/integration/testutils"
)

const numberOfApps = 5 // Number of Applications to be added in a loop

func TestListApp(t *testing.T) {
	username := superAdminUser
	password := superAdminPassword
	otherUsername := subscriber.UserName
	otherPassword := subscriber.Password

	apim := GetDevClient()
	testutils.AddApp(t, apim, username, password)
	testutils.AddApp(t, apim, otherUsername, otherPassword)

	base.SetupEnv(t, apim.GetEnvName(), apim.GetApimURL(), apim.GetTokenURL())
	base.Login(t, apim.GetEnvName(), username, password)
	testutils.ListApps(t, apim.GetEnvName())
}

func TestListAppsDevopsSuperTenantUser(t *testing.T) {
	username := superAdminUser
	password := superAdminPassword

	devopsUsername := devops.UserName
	devopsPassword := devops.Password

	otherUsername := subscriber.UserName
	otherPassword := subscriber.Password

	apim := GetDevClient()
	testutils.AddApp(t, apim, username, password)
	testutils.AddApp(t, apim, otherUsername, otherPassword)

	base.SetupEnv(t, apim.GetEnvName(), apim.GetApimURL(), apim.GetTokenURL())
	base.Login(t, apim.GetEnvName(), devopsUsername, devopsPassword)
	testutils.ListApps(t, apim.GetEnvName())
}

func TestListAppsDevopsTenantUser(t *testing.T) {
	tenantAdminUsername := superAdminUser + "@" + TENANT1
	tenantAdminPassword := superAdminPassword

	tenantDevopsUsername := devops.UserName + "@" + TENANT1
	tenantDevopsPassword := devops.Password

	otherUsername := subscriber.UserName + "@" + TENANT1
	otherPassword := subscriber.Password

	apim := GetDevClient()
	testutils.AddApp(t, apim, tenantAdminUsername, tenantAdminPassword)
	testutils.AddApp(t, apim, otherUsername, otherPassword)

	base.SetupEnv(t, apim.GetEnvName(), apim.GetApimURL(), apim.GetTokenURL())
	base.Login(t, apim.GetEnvName(), tenantDevopsUsername, tenantDevopsPassword)
	testutils.ListApps(t, apim.GetEnvName())
}

//List all the applications in an environment (by specifying the owner)
func TestListAppWithOwner(t *testing.T) {
	username := superAdminUser
	password := superAdminPassword

	apim := GetDevClient()

	for appCount := 0; appCount < 5; appCount++ {
		testutils.AddApp(t, apim, username, password)
	}

	base.SetupEnv(t, apim.GetEnvName(), apim.GetApimURL(), apim.GetTokenURL())
	base.Login(t, apim.GetEnvName(), username, password)

	testutils.ValidateListAppsWithOwner(t, apim.GetEnvName())
}

func TestExportAppNonAdminSuperTenant(t *testing.T) {
	subscriberUserName := subscriber.UserName
	subscriberPassword := subscriber.Password

	dev := GetDevClient()

	app := testutils.AddApp(t, dev, subscriberUserName, subscriberPassword)

	args := &testutils.AppImportExportTestArgs{
		AppOwner:    testutils.Credentials{Username: subscriberUserName, Password: subscriberPassword},
		CtlUser:     testutils.Credentials{Username: subscriberUserName, Password: subscriberPassword},
		Application: app,
		SrcAPIM:     dev,
	}

	testutils.ValidateAppExport(t, args)
}

func TestExportAppNonAdminTenant(t *testing.T) {
	subscriberUserName := subscriber.UserName + "@" + TENANT1
	subscriberPassword := subscriber.Password

	dev := GetDevClient()

	app := testutils.AddApp(t, dev, subscriberUserName, subscriberPassword)

	args := &testutils.AppImportExportTestArgs{
		AppOwner:    testutils.Credentials{Username: subscriberUserName, Password: subscriberPassword},
		CtlUser:     testutils.Credentials{Username: subscriberUserName, Password: subscriberPassword},
		Application: app,
		SrcAPIM:     dev,
	}

	testutils.ValidateAppExport(t, args)
}

func TestExportImportOwnAppAdminSuperTenant(t *testing.T) {
	adminUsername := superAdminUser
	adminPassword := superAdminPassword

	dev := GetDevClient()
	prod := GetProdClient()

	app := testutils.AddApp(t, dev, adminUsername, adminPassword)

	args := &testutils.AppImportExportTestArgs{
		AppOwner:    testutils.Credentials{Username: adminUsername, Password: adminPassword},
		CtlUser:     testutils.Credentials{Username: adminUsername, Password: adminPassword},
		Application: app,
		SrcAPIM:     dev,
		DestAPIM:    prod,
	}

	testutils.ValidateAppExportImportWithPreserveOwner(t, args)
}

//Import an already export App with already generated Keys with --update flag
func TestExportImportOwnAppAdminSuperTenantWithUpdate(t *testing.T) {
	adminUsername := superAdminUser
	adminPassword := superAdminPassword

	dev := GetDevClient()
	prod := GetProdClient()

	app := testutils.AddApp(t, dev, adminUsername, adminPassword)

	args := &testutils.AppImportExportTestArgs{
		AppOwner:    testutils.Credentials{Username: adminUsername, Password: adminPassword},
		CtlUser:     testutils.Credentials{Username: adminUsername, Password: adminPassword},
		Application: app,
		SrcAPIM:     dev,
		DestAPIM:    prod,
	}

	testutils.ValidateAppExportImportWithUpdate(t, args)
}

// Import an already export App with already generated Keys with --update flag (Using a user with Internal/devops role)
func TestExportImportAppDevopsSuperTenantWithUpdate(t *testing.T) {
	adminUsername := superAdminUser
	adminPassword := superAdminPassword

	devopsUsername := devops.UserName
	devopsPassword := devops.Password

	dev := GetDevClient()
	prod := GetProdClient()

	app := testutils.AddApp(t, dev, adminUsername, adminPassword)

	args := &testutils.AppImportExportTestArgs{
		AppOwner:    testutils.Credentials{Username: adminUsername, Password: adminPassword},
		CtlUser:     testutils.Credentials{Username: devopsUsername, Password: devopsPassword},
		Application: app,
		SrcAPIM:     dev,
		DestAPIM:    prod,
	}

	testutils.ValidateAppExportImportWithUpdate(t, args)
}

func TestExportImportOtherAppAdminSuperTenant(t *testing.T) {
	otherUsername := subscriber.UserName
	otherPassword := subscriber.Password
	adminUsername := superAdminUser
	adminPassword := superAdminPassword

	dev := GetDevClient()
	prod := GetProdClient()

	app := testutils.AddApp(t, dev, otherUsername, otherPassword)

	args := &testutils.AppImportExportTestArgs{
		AppOwner:    testutils.Credentials{Username: otherUsername, Password: otherPassword},
		CtlUser:     testutils.Credentials{Username: adminUsername, Password: adminPassword},
		Application: app,
		SrcAPIM:     dev,
		DestAPIM:    prod,
	}

	testutils.ValidateAppExportImportWithPreserveOwner(t, args)
}

// Export an application (created by super tenant user) and import it to another
// environment while preserving the owner by a user with Internal/devops role
func TestExportImportAppDevopsSuperTenant(t *testing.T) {
	devopsUsername := devops.UserName
	devopsPassword := devops.Password

	adminUsername := superAdminUser
	adminPassword := superAdminPassword

	dev := GetDevClient()
	prod := GetProdClient()

	app := testutils.AddApp(t, dev, adminUsername, adminPassword)

	args := &testutils.AppImportExportTestArgs{
		AppOwner:    testutils.Credentials{Username: adminUsername, Password: adminPassword},
		CtlUser:     testutils.Credentials{Username: devopsUsername, Password: devopsPassword},
		Application: app,
		SrcAPIM:     dev,
		DestAPIM:    prod,
	}

	testutils.ValidateAppExportImportWithPreserveOwner(t, args)
}

func TestExportImportOwnAppAdminTenant(t *testing.T) {
	adminUsername := superAdminUser + "@" + TENANT1
	adminPassword := superAdminPassword

	dev := GetDevClient()
	prod := GetProdClient()

	app := testutils.AddApp(t, dev, adminUsername, adminPassword)

	args := &testutils.AppImportExportTestArgs{
		AppOwner:    testutils.Credentials{Username: adminUsername, Password: adminPassword},
		CtlUser:     testutils.Credentials{Username: adminUsername, Password: adminPassword},
		Application: app,
		SrcAPIM:     dev,
		DestAPIM:    prod,
	}

	testutils.ValidateAppExportImportWithPreserveOwner(t, args)
}

func TestExportOtherAppAdminTenant(t *testing.T) {
	otherUsername := subscriber.UserName + "@" + TENANT1
	otherPassword := subscriber.Password
	adminUsername := superAdminUser + "@" + TENANT1
	adminPassword := superAdminPassword

	dev := GetDevClient()
	prod := GetProdClient()

	app := testutils.AddApp(t, dev, otherUsername, otherPassword)

	args := &testutils.AppImportExportTestArgs{
		AppOwner:    testutils.Credentials{Username: otherUsername, Password: otherPassword},
		CtlUser:     testutils.Credentials{Username: adminUsername, Password: adminPassword},
		Application: app,
		SrcAPIM:     dev,
		DestAPIM:    prod,
	}

	testutils.ValidateAppExportImportWithPreserveOwner(t, args)
}

// Export an application (created by tenant user) and import it to another
// environment while preserving the owner by a user with Internal/devops role
func TestExportImportAppDevopsTenant(t *testing.T) {
	tenantDevopsUsername := devops.UserName + "@" + TENANT1
	tenantDevopsPassword := devops.Password

	tenantAdminUsername := superAdminUser + "@" + TENANT1
	tenantAdminPassword := superAdminPassword

	dev := GetDevClient()
	prod := GetProdClient()

	app := testutils.AddApp(t, dev, tenantAdminUsername, tenantAdminPassword)

	args := &testutils.AppImportExportTestArgs{
		AppOwner:    testutils.Credentials{Username: tenantAdminUsername, Password: tenantAdminPassword},
		CtlUser:     testutils.Credentials{Username: tenantDevopsUsername, Password: tenantDevopsPassword},
		Application: app,
		SrcAPIM:     dev,
		DestAPIM:    prod,
	}

	testutils.ValidateAppExportImportWithPreserveOwner(t, args)
}

func TestExportCrossTenantAppAdminTenant(t *testing.T) {
	adminUsername := superAdminUser
	adminPassword := superAdminPassword
	tenantAdminUsername := superAdminUser + "@" + TENANT1
	tenantAdminPassword := superAdminPassword

	dev := GetDevClient()

	app := testutils.AddApp(t, dev, adminUsername, adminPassword)

	args := &testutils.AppImportExportTestArgs{
		AppOwner:    testutils.Credentials{Username: adminUsername, Password: adminPassword},
		CtlUser:     testutils.Credentials{Username: tenantAdminUsername, Password: tenantAdminPassword},
		Application: app,
		SrcAPIM:     dev,
	}

	testutils.ValidateAppExportFailure(t, args)
}

// Export an application (created by a super tenant user) and
// import it to another tenant domain by a user with Internal/devops role
func TestExportCrossTenantAppDevopsTenant(t *testing.T) {
	adminUsername := superAdminUser
	adminPassword := superAdminPassword

	tenantDevopsUsername := devops.UserName + "@" + TENANT1
	tenantDevopsPassword := devops.Password

	dev := GetDevClient()

	app := testutils.AddApp(t, dev, adminUsername, adminPassword)

	args := &testutils.AppImportExportTestArgs{
		AppOwner:    testutils.Credentials{Username: adminUsername, Password: adminPassword},
		CtlUser:     testutils.Credentials{Username: tenantDevopsUsername, Password: tenantDevopsPassword},
		Application: app,
		SrcAPIM:     dev,
	}

	testutils.ValidateAppExportFailure(t, args)
}

// TODO: Secondary user store test cases, need to enabled when later on when secondary user store creation is automated
/*
func TestExportAppSecondaryUserStoreAdminSuperTenant(t *testing.T) {
	username := "SECOND.COM/super"
	password := "admin"

	name := "DefaultApplication"
	owner := "SECOND.COM/super"

	base.SetupEnv(t, devEnv, devApim, devTokenEP)
	base.Login(t, devEnv, username, password)
validateAppExportImportWithPreserveOwner
	exportApp(t, name, owner, devEnv)

	assert.True(t, base.IsApplicationArchiveExists(devAppExportPath, name, owner))
}

func TestExportAppSecondaryUserStoreAdminSuperTenantLowerCase(t *testing.T) {
	username := "second.com/super"
	password := "admin"

	name := "DefaultApplication"
	owner := "second.com/super"

	base.SetupEnv(t, devEnv, devApim, devTokenEP)
	base.Login(t, devEnv, username, password)

	exportApp(t, name, owner, devEnv)

	assert.True(t, base.IsApplicationArchiveExists(devAppExportPath, name, owner))
}
*/

//Delete an Application as a super tenant admin
func TestDeleteAppSuperTenantUser(t *testing.T) {
	adminUsername := superAdminUser
	adminPassword := superAdminPassword

	dev := GetDevClient()

	var application *apim.Application
	for appCount := 0; appCount <= numberOfApps; appCount++ {
		application = testutils.AddApp(t, dev, adminUsername, adminPassword)
	}

	// This will be the Application that will be deleted by apictl, so no need to do cleaning
	application = testutils.AddApplicationWithoutCleaning(t, dev, adminUsername, adminPassword)

	args := &testutils.AppImportExportTestArgs{
		CtlUser:     testutils.Credentials{Username: superAdminUser, Password: superAdminPassword},
		Application: application,
		SrcAPIM:     dev,
	}

	testutils.ValidateAppDelete(t, args)
}


// Export an application with space in application name  and import it to another  to check whether the url
// encoding is working properly
func TestExportImportOwnAppWithSpaceInAppName(t *testing.T) {
	for _, user := range testCaseUsers {
		t.Run(user.Description, func(t *testing.T) {
			dev := GetDevClient()
			prod := GetProdClient()

			app := testutils.AddAppWithSpaceInAppName(t, dev, user.CtlUser.Username, user.CtlUser.Password)

			args := &testutils.AppImportExportTestArgs{
				AppOwner:      testutils.Credentials{Username: user.CtlUser.Username, Password: user.CtlUser.Password},
				CtlUser:       testutils.Credentials{Username: user.CtlUser.Username, Password: user.CtlUser.Password},
				Application:   app,
				SrcAPIM:       dev,
				DestAPIM:      prod,
				PreserveOwner: true,
			}

			testutils.ValidateAppExportImport(t, args, true)
		})
	}
}
