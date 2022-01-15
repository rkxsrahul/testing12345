package web

import (
	"log"
	"os"
	"strconv"
	"testing"

	"git.xenonstack.com/util/continuous-security-backend/config"
	"git.xenonstack.com/util/continuous-security-backend/src/database"
	"git.xenonstack.com/util/continuous-security-backend/src/nats"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func ReconnectDatabase() {
	db, err := gorm.Open("sqlite3", os.Getenv("HOME")+"/account-testing.db")
	if err != nil {
		log.Println(err)
		log.Println("Exit")
		os.Exit(1)
	}
	config.DB = db
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	os.Remove(os.Getenv("HOME") + "/account-testing.db")
	db, err := gorm.Open("sqlite3", os.Getenv("HOME")+"/account-testing.db")
	if err != nil {
		log.Println(err)
		log.Println("Exit")
		os.Exit(1)
	}
	config.ConfigurationWithToml("../../example.toml")
	go nats.InitConnection()
	config.DB = db
	//create table
	database.CreateDBTablesIfNotExists()

	data := database.RequestInfo{}
	data.ID = 1
	data.Email = "test@xenonstack.com"
	data.Workspace = "test"

	db.Create(&data)

	website := database.ScanResult{}
	website.ID = 1
	website.Method = "Website Security"
	website.UUID = "dwoidjfdefjejwf"

	err = db.Create(&website).Error
	log.Println(err)

	list := database.RequestInfo{}
	list.Email = "test@xenonstack.com"
	list.ID = 888
	list.Name = "test"
	list.Workspace = "test"
	list.RepoLang = "Node Scan"
	list.UUID = "ig3wih3ii3nhi3iioi"

	list4 := database.RequestInfo{}
	list4.Email = "test@xenonstack.com"
	list4.ID = 856
	list4.Name = "test"
	list4.Workspace = "test"
	list4.RepoLang = "Node Scan"
	list4.UUID = "ig3wih3ii33333333"

	db.Create(&list4)

	list2 := database.RequestInfo{}
	list2.Email = "test@xenonstack.com"
	list2.ID = 2
	list2.Name = "test2"
	list2.Workspace = "test"
	list2.UUID = "dwoidjfdefjejwfhhohjhoi"
	list2.RepoLang = ""

	db.Create(&list)
	err = db.Create(&list2).Error
	log.Println(err)

	existInfo := database.RequestInfo{}
	existInfo.UUID = "testewihgfiueiu"
	existInfo.Email = "test22@xenonstack.com"
	existInfo.URL = "https://www.xenonstack.com"

	db.Create(&existInfo)

	for i := 0; i < 15; i++ {
		website := database.ScanResult{}
		website.ID = i + 22
		website.Method = "Website Security"
		website.UUID = "dwoidjfdefjejwfhhohjhoi"
		website.CommandName = "ororog" + strconv.Itoa(i)
		log.Println(website.CommandName)
		err := db.Create(&website).Error
		log.Println(err)
	}
	for i := 0; i < 3; i++ {
		website := database.ScanResult{}
		website.ID = i + 88
		website.Method = "Email Security"
		website.UUID = "dwoidjfdefjejwfhhohjhoi"
		website.CommandName = "hhhrorog" + strconv.Itoa(i)
		log.Println(i)
		err := db.Create(&website).Error
		log.Println(err)
	}
	for i := 0; i < 2; i++ {
		website := database.ScanResult{}
		website.ID = i + 44
		website.Method = "Network Security"
		website.UUID = "dwoidjfdefjejwfhhohjhoi"
		website.CommandName = "hhhjjorog" + strconv.Itoa(i)
		log.Println(i)
		err := db.Create(&website).Error
		log.Println(err)
	}
	for i := 0; i < 7; i++ {
		website := database.ScanResult{}
		website.ID = i + 55
		website.Method = "HTTP Security Headers"
		website.UUID = "dwoidjfdefjejwfhhohjhoi"
		website.CommandName = "hhhkkkorog" + strconv.Itoa(i)
		log.Println(i)
		db.Create(&website)
	}

	website22 := database.ScanResult{}
	website22.ID = 77
	website22.UUID = "ig3wih3ii3nhi3iioi"
	website22.CommandName = "hhhkkkorog"
	website22.Result = `{"Vulnerabilities":[{"CVSS":{"nvd":{"V2Score":5,"V2Vector":"AV:N/AC:L/Au:N/C:P/I:N/A:N","V3Score":9.8,"V3Vector":"CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H"},"redhat":{"V3Score":3.7,"V3Vector":"CVSS:3.0/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:N/A:N"}},"CweIDs":["CWE-331"],"Description":"Eran Hammer cryptiles version 4.1.1 earlier contains a CWE-331: Insufficient Entropy vulnerability in randomDigits() method that can result in An attacker is more likely to be able to brute force something that was supposed to be random.. This attack appear to be exploitable via Depends upon the calling application.. This vulnerability appears to have been fixed in 4.1.2.","FixedVersion":"4.1.2","InstalledVersion":"2.0.5","LastModifiedDate":"2018-09-10T16:07:00Z","Layer":{},"PkgName":"cryptiles","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2018-1000620","PublishedDate":"2018-07-09T20:29:00Z","References":["https://github.com/advisories/GHSA-rq8g-5pc5-wrhr","https://github.com/hapijs/cryptiles/issues/34","https://github.com/nodejs/security-wg/blob/master/vuln/npm/476.json","https://nvd.nist.gov/vuln/detail/CVE-2018-1000620","https://www.npmjs.com/advisories/1464","https://www.npmjs.com/advisories/720"],"Severity":"CRITICAL","SeveritySource":"nodejs-security-wg","Title":"nodejs-cryptiles: Insecure randomness causes the randomDigits() function returns a pseudo-random data string biased to certain digits","VulnerabilityID":"CVE-2018-1000620"},{"CVSS":{"nvd":{"V2Score":6.5,"V2Vector":"AV:N/AC:L/Au:S/C:P/I:P/A:P","V3Score":8.8,"V3Vector":"CVSS:3.0/AV:N/AC:L/PR:L/UI:N/S:U/C:H/I:H/A:H"},"redhat":{"V3Score":2.9,"V3Vector":"CVSS:3.0/AV:L/AC:H/PR:N/UI:N/S:U/C:N/I:N/A:L"}},"CweIDs":["CWE-471"],"Description":"hoek node module before 4.2.0 and 5.0.x before 5.0.3 suffers from a Modification of Assumed-Immutable Data (MAID) vulnerability via 'merge' and 'applyToDefaults' functions, which allows a malicious user to modify the prototype of \"Object\" via __proto__, causing the addition or modification of an existing property that will exist on all objects.","FixedVersion":"5.0.3, 4.2.1","InstalledVersion":"2.16.3","LastModifiedDate":"2019-10-09T23:40:00Z","Layer":{},"PkgName":"hoek","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2018-3728","PublishedDate":"2018-03-30T19:29:00Z","References":["http://www.securityfocus.com/bid/103108","https://access.redhat.com/errata/RHSA-2018:1263","https://access.redhat.com/errata/RHSA-2018:1264","https://github.com/advisories/GHSA-jp4x-w63m-7wgm","https://github.com/hapijs/hoek/commit/32ed5c9413321fbc37da5ca81a7cbab693786dee","https://hackerone.com/reports/310439","https://nodesecurity.io/advisories/566","https://nvd.nist.gov/vuln/detail/CVE-2018-3728","https://snyk.io/vuln/npm:hoek:20180212","https://www.npmjs.com/advisories/566"],"Severity":"ELSE","SeveritySource":"nodejs-security-wg","Title":"hoek: Prototype pollution in utilities function","VulnerabilityID":"CVE-2018-3728"},{"CVSS":{"nvd":{"V2Score":6.4,"V2Vector":"AV:N/AC:L/Au:N/C:N/I:P/A:P","V3Score":9.1,"V3Vector":"CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:H/A:H"},"redhat":{"V3Score":9.1,"V3Vector":"CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:H/A:H"}},"Description":"Versions of lodash lower than 4.17.12 are vulnerable to Prototype Pollution. The function defaultsDeep could be tricked into adding or modifying properties of Object.prototype using a constructor payload.","FixedVersion":"4.17.12","InstalledVersion":"4.17.4","LastModifiedDate":"2021-03-16T13:57:00Z","Layer":{},"PkgName":"lodash","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2019-10744","PublishedDate":"2019-07-26T00:15:00Z","References":["https://access.redhat.com/errata/RHSA-2019:3024","https://github.com/advisories/GHSA-jf85-cpcp-j695","https://github.com/lodash/lodash/pull/4336","https://nvd.nist.gov/vuln/detail/CVE-2019-10744","https://security.netapp.com/advisory/ntap-20191004-0005/","https://snyk.io/vuln/SNYK-JS-LODASH-450202","https://support.f5.com/csp/article/K47105354?utm_source=f5support\u0026amp;utm_medium=RSS","https://www.npmjs.com/advisories/1065","https://www.oracle.com/security-alerts/cpujan2021.html","https://www.oracle.com/security-alerts/cpuoct2020.html"],"Severity":"CRITICAL","SeveritySource":"nvd","Title":"nodejs-lodash: prototype pollution in defaultsDeep function leading to modifying properties","VulnerabilityID":"CVE-2019-10744"},{"CVSS":{"nvd":{"V2Score":6.8,"V2Vector":"AV:N/AC:M/Au:N/C:P/I:P/A:P","V3Score":5.6,"V3Vector":"CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L"},"redhat":{"V3Score":5.6,"V3Vector":"CVSS:3.0/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L"}},"Description":"A prototype pollution vulnerability was found in lodash \u003c4.17.11 where the functions merge, mergeWith, and defaultsDeep can be tricked into adding or modifying properties of Object.prototype.","FixedVersion":"4.17.11","InstalledVersion":"4.17.4","LastModifiedDate":"2020-09-18T16:38:00Z","Layer":{},"PkgName":"lodash","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2018-16487","PublishedDate":"2019-02-01T18:29:00Z","References":["https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-16487","https://github.com/advisories/GHSA-4xc9-xhrj-v574","https://hackerone.com/reports/380873","https://nvd.nist.gov/vuln/detail/CVE-2018-16487","https://security.netapp.com/advisory/ntap-20190919-0004/","https://www.npmjs.com/advisories/782"],"Severity":"HIGH","SeveritySource":"nodejs-security-wg","Title":"lodash: Prototype pollution in utilities function","VulnerabilityID":"CVE-2018-16487"},{"CVSS":{"nvd":{"V2Score":5.8,"V2Vector":"AV:N/AC:M/Au:N/C:N/I:P/A:P","V3Score":7.4,"V3Vector":"CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:N/I:H/A:H"},"redhat":{"V3Score":7.4,"V3Vector":"CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:N/I:H/A:H"}},"CweIDs":["CWE-1321"],"Description":"Prototype pollution attack when using _.zipObjectDeep in lodash before 4.17.20.","FixedVersion":"4.17.19","InstalledVersion":"4.17.4","LastModifiedDate":"2021-12-02T22:14:00Z","Layer":{},"PkgName":"lodash","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2020-8203","PublishedDate":"2020-07-15T17:15:00Z","References":["https://github.com/advisories/GHSA-p6mc-m468-83gw","https://github.com/lodash/lodash/commit/c84fe82760fb2d3e03a63379b297a1cc1a2fce12","https://github.com/lodash/lodash/issues/4744","https://github.com/lodash/lodash/issues/4874","https://hackerone.com/reports/712065","https://nvd.nist.gov/vuln/detail/CVE-2020-8203","https://security.netapp.com/advisory/ntap-20200724-0006/","https://www.npmjs.com/advisories/1523","https://www.oracle.com//security-alerts/cpujul2021.html","https://www.oracle.com/security-alerts/cpuApr2021.html","https://www.oracle.com/security-alerts/cpuoct2021.html"],"Severity":"HIGH","SeveritySource":"nvd","Title":"nodejs-lodash: prototype pollution in zipObjectDeep function","VulnerabilityID":"CVE-2020-8203"},{"CVSS":{"nvd":{"V2Score":6.5,"V2Vector":"AV:N/AC:L/Au:S/C:P/I:P/A:P","V3Score":7.2,"V3Vector":"CVSS:3.1/AV:N/AC:L/PR:H/UI:N/S:U/C:H/I:H/A:H"},"redhat":{"V3Score":7.2,"V3Vector":"CVSS:3.1/AV:N/AC:L/PR:H/UI:N/S:U/C:H/I:H/A:H"}},"CweIDs":["CWE-77"],"Description":"Lodash versions prior to 4.17.21 are vulnerable to Command Injection via the template function.","FixedVersion":"4.17.21","InstalledVersion":"4.17.4","LastModifiedDate":"2021-12-07T20:52:00Z","Layer":{},"PkgName":"lodash","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2021-23337","PublishedDate":"2021-02-15T13:15:00Z","References":["https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-23337","https://github.com/advisories/GHSA-35jh-r3h4-6jhm","https://github.com/lodash/lodash/blob/ddfd9b11a0126db2302cb70ec9973b66baec0975/lodash.js#L14851","https://github.com/lodash/lodash/blob/ddfd9b11a0126db2302cb70ec9973b66baec0975/lodash.js%23L14851","https://github.com/lodash/lodash/commit/3469357cff396a26c363f8c1b5a91dde28ba4b1c","https://nvd.nist.gov/vuln/detail/CVE-2021-23337","https://security.netapp.com/advisory/ntap-20210312-0006/","https://snyk.io/vuln/SNYK-JAVA-ORGFUJIONWEBJARS-1074932","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARS-1074930","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARSBOWER-1074928","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARSBOWERGITHUBLODASH-1074931","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARSNPM-1074929","https://snyk.io/vuln/SNYK-JS-LODASH-1040724","https://www.oracle.com//security-alerts/cpujul2021.html","https://www.oracle.com/security-alerts/cpuoct2021.html"],"Severity":"HIGH","SeveritySource":"nvd","Title":"nodejs-lodash: command injection via template","VulnerabilityID":"CVE-2021-23337"},{"CVSS":{"nvd":{"V2Score":4,"V2Vector":"AV:N/AC:L/Au:S/C:N/I:N/A:P","V3Score":6.5,"V3Vector":"CVSS:3.1/AV:N/AC:L/PR:L/UI:N/S:U/C:N/I:N/A:H"},"redhat":{"V3Score":4.4,"V3Vector":"CVSS:3.0/AV:N/AC:H/PR:H/UI:N/S:U/C:N/I:N/A:H"}},"CweIDs":["CWE-770"],"Description":"lodash prior to 4.17.11 is affected by: CWE-400: Uncontrolled Resource Consumption. The impact is: Denial of service. The component is: Date handler. The attack vector is: Attacker provides very long strings, which the library attempts to match using a regular expression. The fixed version is: 4.17.11.","FixedVersion":"4.17.11","InstalledVersion":"4.17.4","LastModifiedDate":"2020-09-30T13:40:00Z","Layer":{},"PkgName":"lodash","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2019-1010266","PublishedDate":"2019-07-17T21:15:00Z","References":["https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2019-1010266","https://github.com/advisories/GHSA-x5rq-j2xg-h7qm","https://github.com/lodash/lodash/commit/5c08f18d365b64063bfbfa686cbb97cdd6267347","https://github.com/lodash/lodash/issues/3359","https://github.com/lodash/lodash/wiki/Changelog","https://nvd.nist.gov/vuln/detail/CVE-2019-1010266","https://security.netapp.com/advisory/ntap-20190919-0004/","https://snyk.io/vuln/SNYK-JS-LODASH-73639"],"Severity":"MEDIUM","SeveritySource":"nvd","Title":"lodash: uncontrolled resource consumption in Data handler causing denial of service","VulnerabilityID":"CVE-2019-1010266"},{"CVSS":{"nvd":{"V2Score":5,"V2Vector":"AV:N/AC:L/Au:N/C:N/I:N/A:P","V3Score":5.3,"V3Vector":"CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:L"},"redhat":{"V3Score":5.3,"V3Vector":"CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:L"}},"Description":"Lodash versions prior to 4.17.21 are vulnerable to Regular Expression Denial of Service (ReDoS) via the toNumber, trim and trimEnd functions.","FixedVersion":"4.17.21","InstalledVersion":"4.17.4","LastModifiedDate":"2021-12-10T18:04:00Z","Layer":{},"PkgName":"lodash","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2020-28500","PublishedDate":"2021-02-15T11:15:00Z","References":["https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-28500","https://github.com/advisories/GHSA-29mw-wpgm-hmr9","https://github.com/lodash/lodash/blob/npm/trimEnd.js#L8","https://github.com/lodash/lodash/blob/npm/trimEnd.js%23L8","https://github.com/lodash/lodash/pull/5065","https://github.com/lodash/lodash/pull/5065/commits/02906b8191d3c100c193fe6f7b27d1c40f200bb7","https://nvd.nist.gov/vuln/detail/CVE-2020-28500","https://security.netapp.com/advisory/ntap-20210312-0006/","https://snyk.io/vuln/SNYK-JAVA-ORGFUJIONWEBJARS-1074896","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARS-1074894","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARSBOWER-1074892","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARSBOWERGITHUBLODASH-1074895","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARSNPM-1074893","https://snyk.io/vuln/SNYK-JS-LODASH-1018905","https://www.oracle.com//security-alerts/cpujul2021.html","https://www.oracle.com/security-alerts/cpuoct2021.html"],"Severity":"MEDIUM","SeveritySource":"nvd","Title":"nodejs-lodash: ReDoS via the toNumber, trim and trimEnd functions","VulnerabilityID":"CVE-2020-28500"},{"CVSS":{"nvd":{"V2Score":4,"V2Vector":"AV:N/AC:L/Au:S/C:N/I:P/A:N","V3Score":6.5,"V3Vector":"CVSS:3.0/AV:N/AC:L/PR:L/UI:N/S:U/C:N/I:H/A:N"},"redhat":{"V3Score":2.9,"V3Vector":"CVSS:3.0/AV:L/AC:H/PR:N/UI:N/S:U/C:N/I:N/A:L"}},"Description":"lodash node module before 4.17.5 suffers from a Modification of Assumed-Immutable Data (MAID) vulnerability via defaultsDeep, merge, and mergeWith functions, which allows a malicious user to modify the prototype of \"Object\" via __proto__, causing the addition or modification of an existing property that will exist on all objects.","FixedVersion":"4.17.5","InstalledVersion":"4.17.4","LastModifiedDate":"2019-10-03T00:03:00Z","Layer":{},"PkgName":"lodash","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2018-3721","PublishedDate":"2018-06-07T02:29:00Z","References":["https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-3721","https://github.com/advisories/GHSA-fvqr-27wr-82fm","https://github.com/lodash/lodash/commit/d8e069cc3410082e44eb18fcf8e7f3d08ebe1d4a","https://hackerone.com/reports/310443","https://nvd.nist.gov/vuln/detail/CVE-2018-3721","https://security.netapp.com/advisory/ntap-20190919-0004/","https://snyk.io/vuln/npm:lodash:20180130","https://www.npmjs.com/advisories/577"],"Severity":"LOW","SeveritySource":"nodejs-security-wg","Title":"lodash: Prototype pollution in utilities function","VulnerabilityID":"CVE-2018-3721"},{"CVSS":{"nvd":{"V2Score":6.8,"V2Vector":"AV:N/AC:M/Au:N/C:P/I:P/A:P","V3Score":5.6,"V3Vector":"CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L"},"redhat":{"V3Score":5.6,"V3Vector":"CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L"}},"CweIDs":["CWE-20"],"Description":"minimist before 1.2.2 could be tricked into adding or modifying properties of Object.prototype using a \"constructor\" or \"__proto__\" payload.","FixedVersion":"1.2.3, 0.2.1","InstalledVersion":"0.0.8","LastModifiedDate":"2021-07-21T11:39:00Z","Layer":{},"PkgName":"minimist","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2020-7598","PublishedDate":"2020-03-11T23:15:00Z","References":["http://lists.opensuse.org/opensuse-security-announce/2020-06/msg00024.html","https://github.com/advisories/GHSA-vh95-rmgr-6w4m","https://github.com/substack/minimist/commit/38a4d1caead72ef99e824bb420a2528eec03d9ab","https://github.com/substack/minimist/commit/4cf1354839cb972e38496d35e12f806eea92c11f#diff-a1e0ee62c91705696ddb71aa30ad4f95","https://github.com/substack/minimist/commit/63e7ed05aa4b1889ec2f3b196426db4500cbda94","https://linux.oracle.com/cve/CVE-2020-7598.html","https://linux.oracle.com/errata/ELSA-2020-2852.html","https://nvd.nist.gov/vuln/detail/CVE-2020-7598","https://snyk.io/vuln/SNYK-JS-MINIMIST-559764","https://www.npmjs.com/advisories/1179"],"Severity":"MEDIUM","SeveritySource":"nvd","Title":"nodejs-minimist: prototype pollution allows adding or modifying properties of Object.prototype using a constructor or __proto__ payload","VulnerabilityID":"CVE-2020-7598"},{"CVSS":{"nvd":{"V2Score":6.8,"V2Vector":"AV:N/AC:M/Au:N/C:P/I:P/A:P","V3Score":5.6,"V3Vector":"CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L"},"redhat":{"V3Score":5.6,"V3Vector":"CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L"}},"CweIDs":["CWE-20"],"Description":"minimist before 1.2.2 could be tricked into adding or modifying properties of Object.prototype using a \"constructor\" or \"__proto__\" payload.","FixedVersion":"1.2.3, 0.2.1","InstalledVersion":"1.2.0","LastModifiedDate":"2021-07-21T11:39:00Z","Layer":{},"PkgName":"minimist","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2020-7598","PublishedDate":"2020-03-11T23:15:00Z","References":["http://lists.opensuse.org/opensuse-security-announce/2020-06/msg00024.html","https://github.com/advisories/GHSA-vh95-rmgr-6w4m","https://github.com/substack/minimist/commit/38a4d1caead72ef99e824bb420a2528eec03d9ab","https://github.com/substack/minimist/commit/4cf1354839cb972e38496d35e12f806eea92c11f#diff-a1e0ee62c91705696ddb71aa30ad4f95","https://github.com/substack/minimist/commit/63e7ed05aa4b1889ec2f3b196426db4500cbda94","https://linux.oracle.com/cve/CVE-2020-7598.html","https://linux.oracle.com/errata/ELSA-2020-2852.html","https://nvd.nist.gov/vuln/detail/CVE-2020-7598","https://snyk.io/vuln/SNYK-JS-MINIMIST-559764","https://www.npmjs.com/advisories/1179"],"Severity":"MEDIUM","SeveritySource":"nvd","Title":"nodejs-minimist: prototype pollution allows adding or modifying properties of Object.prototype using a constructor or __proto__ payload","VulnerabilityID":"CVE-2020-7598"}]}`
	website22.Status = true
	db.Create(&website22)

	info := database.ScanResult{}
	info.UUID = "testkrjfnrjnfjrnnf1"
	info.Method = "Website Security"
	info.Result = `"impact":"HIGH"`
	info.CommandName = "oehfoejfjeoif"
	db.Create(&info)
	info2 := database.ScanResult{}
	info2.UUID = "testkrjfnrjnfjrnnf1"
	info2.Method = "Website Security"
	info2.Result = `"impact":"MEDIUM"`
	info2.CommandName = "oehfoejfjeoi2"
	db.Create(&info2)
	info3 := database.ScanResult{}
	info3.UUID = "testkrjfnrjnfjrnnf1"
	info3.Method = "Website Security"
	info3.Result = `"impact":"LOW"`
	info3.CommandName = "oehfoejfjeoi3"
	db.Create(&info3)

	info4 := database.ScanResult{}
	info4.UUID = "testkrjfnrjnfjrnnf1"
	info4.Method = "Email Security"
	info4.Result = `"impact":"HIGH"`
	info4.CommandName = "oehfoejfjeoi66"
	db.Create(&info4)
	info5 := database.ScanResult{}
	info5.UUID = "testkrjfnrjnfjrnnf1"
	info5.Method = "Email Security"
	info5.Result = `"impact":"MEDIUM"`
	info5.CommandName = "oehfoejfjeoi266"
	db.Create(&info5)
	info6 := database.ScanResult{}
	info6.UUID = "testkrjfnrjnfjrnnf1"
	info6.Method = "Email Security"
	info6.Result = `"impact":"LOW"`
	info6.CommandName = "oehfoejfjeoi366"
	db.Create(&info6)

	info7 := database.ScanResult{}
	info7.UUID = "testkrjfnrjnfjrnnf1"
	info7.Method = "Network Security"
	info7.Result = `"impact":"HIGH"`
	info7.CommandName = "oehfoejfjeoi667"
	db.Create(&info7)
	info8 := database.ScanResult{}
	info8.UUID = "testkrjfnrjnfjrnnf1"
	info8.Method = "Network Security"
	info8.Result = `"impact":"MEDIUM"`
	info8.CommandName = "oehfoejfjeoi2667"
	db.Create(&info8)
	info9 := database.ScanResult{}
	info9.UUID = "testkrjfnrjnfjrnnf1"
	info9.Method = "Network Security"
	info9.Result = `"impact":"LOW"`
	info9.CommandName = "oehfoejfjeoi3667"
	db.Create(&info9)

	info10 := database.ScanResult{}
	info10.UUID = "testkrjfnrjnfjrnnf1"
	info10.Method = "HTTP Security Headers"
	info10.Result = `"impact":"HIGH"`
	info10.CommandName = "oehfoejfjeoi6671"
	db.Create(&info10)
	info11 := database.ScanResult{}
	info11.UUID = "testkrjfnrjnfjrnnf1"
	info11.Method = "HTTP Security Headers"
	info11.Result = `"impact":"MEDIUM"`
	info11.CommandName = "oehfoejfjeoi26671"
	db.Create(&info11)
	info12 := database.ScanResult{}
	info12.UUID = "testkrjfnrjnfjrnnf1"
	info12.Method = "HTTP Security Headers"
	info12.Result = `"impact":"LOW"`
	info12.CommandName = "oehfoejfjeoi36671"
	db.Create(&info12)

	info13 := database.ScanResult{}
	info13.UUID = "testkrjfnrjnfjrnnf12"
	info13.Method = "HTTP Security Headers"
	info13.CommandName = "oehfoejfjeoi366712"
	db.Create(&info13)

	website223 := database.ScanResult{}
	website223.ID = 7733
	website223.UUID = "ig3wih3ii3nhi3iioi3"
	website223.CommandName = "hhhkkkorog3jojho"
	website223.Result = `{"Vulnerabilities":[{"CVSS":{"nvd":{"V2Score":5,"V2Vector":"AV:N/AC:L/Au:N/C:P/I:N/A:N","V3Score":9.8,"V3Vector":"CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H"},"redhat":{"V3Score":3.7,"V3Vector":"CVSS:3.0/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:N/A:N"}},"CweIDs":["CWE-331"],"Description":"Eran Hammer cryptiles version 4.1.1 earlier contains a CWE-331: Insufficient Entropy vulnerability in randomDigits() method that can result in An attacker is more likely to be able to brute force something that was supposed to be random.. This attack appear to be exploitable via Depends upon the calling application.. This vulnerability appears to have been fixed in 4.1.2.","FixedVersion":"4.1.2","InstalledVersion":"2.0.5","LastModifiedDate":"2018-09-10T16:07:00Z","Layer":{},"PkgName":"cryptiles","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2018-1000620","PublishedDate":"2018-07-09T20:29:00Z","References":["https://github.com/advisories/GHSA-rq8g-5pc5-wrhr","https://github.com/hapijs/cryptiles/issues/34","https://github.com/nodejs/security-wg/blob/master/vuln/npm/476.json","https://nvd.nist.gov/vuln/detail/CVE-2018-1000620","https://www.npmjs.com/advisories/1464","https://www.npmjs.com/advisories/720"],"Severity":"CRITICAL","SeveritySource":"nodejs-security-wg","Title":"nodejs-cryptiles: Insecure randomness causes the randomDigits() function returns a pseudo-random data string biased to certain digits","VulnerabilityID":"CVE-2018-1000620"},{"CVSS":{"nvd":{"V2Score":6.5,"V2Vector":"AV:N/AC:L/Au:S/C:P/I:P/A:P","V3Score":8.8,"V3Vector":"CVSS:3.0/AV:N/AC:L/PR:L/UI:N/S:U/C:H/I:H/A:H"},"redhat":{"V3Score":2.9,"V3Vector":"CVSS:3.0/AV:L/AC:H/PR:N/UI:N/S:U/C:N/I:N/A:L"}},"CweIDs":["CWE-471"],"Description":"hoek node module before 4.2.0 and 5.0.x before 5.0.3 suffers from a Modification of Assumed-Immutable Data (MAID) vulnerability via 'merge' and 'applyToDefaults' functions, which allows a malicious user to modify the prototype of \"Object\" via __proto__, causing the addition or modification of an existing property that will exist on all objects.","FixedVersion":"5.0.3, 4.2.1","InstalledVersion":"2.16.3","LastModifiedDate":"2019-10-09T23:40:00Z","Layer":{},"PkgName":"hoek","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2018-3728","PublishedDate":"2018-03-30T19:29:00Z","References":["http://www.securityfocus.com/bid/103108","https://access.redhat.com/errata/RHSA-2018:1263","https://access.redhat.com/errata/RHSA-2018:1264","https://github.com/advisories/GHSA-jp4x-w63m-7wgm","https://github.com/hapijs/hoek/commit/32ed5c9413321fbc37da5ca81a7cbab693786dee","https://hackerone.com/reports/310439","https://nodesecurity.io/advisories/566","https://nvd.nist.gov/vuln/detail/CVE-2018-3728","https://snyk.io/vuln/npm:hoek:20180212","https://www.npmjs.com/advisories/566"],"Severity":"LOW","SeveritySource":"nodejs-security-wg","Title":"hoek: Prototype pollution in utilities function","VulnerabilityID":"CVE-2018-3728"},{"CVSS":{"nvd":{"V2Score":6.4,"V2Vector":"AV:N/AC:L/Au:N/C:N/I:P/A:P","V3Score":9.1,"V3Vector":"CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:H/A:H"},"redhat":{"V3Score":9.1,"V3Vector":"CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:H/A:H"}},"Description":"Versions of lodash lower than 4.17.12 are vulnerable to Prototype Pollution. The function defaultsDeep could be tricked into adding or modifying properties of Object.prototype using a constructor payload.","FixedVersion":"4.17.12","InstalledVersion":"4.17.4","LastModifiedDate":"2021-03-16T13:57:00Z","Layer":{},"PkgName":"lodash","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2019-10744","PublishedDate":"2019-07-26T00:15:00Z","References":["https://access.redhat.com/errata/RHSA-2019:3024","https://github.com/advisories/GHSA-jf85-cpcp-j695","https://github.com/lodash/lodash/pull/4336","https://nvd.nist.gov/vuln/detail/CVE-2019-10744","https://security.netapp.com/advisory/ntap-20191004-0005/","https://snyk.io/vuln/SNYK-JS-LODASH-450202","https://support.f5.com/csp/article/K47105354?utm_source=f5support\u0026amp;utm_medium=RSS","https://www.npmjs.com/advisories/1065","https://www.oracle.com/security-alerts/cpujan2021.html","https://www.oracle.com/security-alerts/cpuoct2020.html"],"Severity":"CRITICAL","SeveritySource":"nvd","Title":"nodejs-lodash: prototype pollution in defaultsDeep function leading to modifying properties","VulnerabilityID":"CVE-2019-10744"},{"CVSS":{"nvd":{"V2Score":6.8,"V2Vector":"AV:N/AC:M/Au:N/C:P/I:P/A:P","V3Score":5.6,"V3Vector":"CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L"},"redhat":{"V3Score":5.6,"V3Vector":"CVSS:3.0/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L"}},"Description":"A prototype pollution vulnerability was found in lodash \u003c4.17.11 where the functions merge, mergeWith, and defaultsDeep can be tricked into adding or modifying properties of Object.prototype.","FixedVersion":"4.17.11","InstalledVersion":"4.17.4","LastModifiedDate":"2020-09-18T16:38:00Z","Layer":{},"PkgName":"lodash","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2018-16487","PublishedDate":"2019-02-01T18:29:00Z","References":["https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-16487","https://github.com/advisories/GHSA-4xc9-xhrj-v574","https://hackerone.com/reports/380873","https://nvd.nist.gov/vuln/detail/CVE-2018-16487","https://security.netapp.com/advisory/ntap-20190919-0004/","https://www.npmjs.com/advisories/782"],"Severity":"HIGH","SeveritySource":"nodejs-security-wg","Title":"lodash: Prototype pollution in utilities function","VulnerabilityID":"CVE-2018-16487"},{"CVSS":{"nvd":{"V2Score":5.8,"V2Vector":"AV:N/AC:M/Au:N/C:N/I:P/A:P","V3Score":7.4,"V3Vector":"CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:N/I:H/A:H"},"redhat":{"V3Score":7.4,"V3Vector":"CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:N/I:H/A:H"}},"CweIDs":["CWE-1321"],"Description":"Prototype pollution attack when using _.zipObjectDeep in lodash before 4.17.20.","FixedVersion":"4.17.19","InstalledVersion":"4.17.4","LastModifiedDate":"2021-12-02T22:14:00Z","Layer":{},"PkgName":"lodash","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2020-8203","PublishedDate":"2020-07-15T17:15:00Z","References":["https://github.com/advisories/GHSA-p6mc-m468-83gw","https://github.com/lodash/lodash/commit/c84fe82760fb2d3e03a63379b297a1cc1a2fce12","https://github.com/lodash/lodash/issues/4744","https://github.com/lodash/lodash/issues/4874","https://hackerone.com/reports/712065","https://nvd.nist.gov/vuln/detail/CVE-2020-8203","https://security.netapp.com/advisory/ntap-20200724-0006/","https://www.npmjs.com/advisories/1523","https://www.oracle.com//security-alerts/cpujul2021.html","https://www.oracle.com/security-alerts/cpuApr2021.html","https://www.oracle.com/security-alerts/cpuoct2021.html"],"Severity":"HIGH","SeveritySource":"nvd","Title":"nodejs-lodash: prototype pollution in zipObjectDeep function","VulnerabilityID":"CVE-2020-8203"},{"CVSS":{"nvd":{"V2Score":6.5,"V2Vector":"AV:N/AC:L/Au:S/C:P/I:P/A:P","V3Score":7.2,"V3Vector":"CVSS:3.1/AV:N/AC:L/PR:H/UI:N/S:U/C:H/I:H/A:H"},"redhat":{"V3Score":7.2,"V3Vector":"CVSS:3.1/AV:N/AC:L/PR:H/UI:N/S:U/C:H/I:H/A:H"}},"CweIDs":["CWE-77"],"Description":"Lodash versions prior to 4.17.21 are vulnerable to Command Injection via the template function.","FixedVersion":"4.17.21","InstalledVersion":"4.17.4","LastModifiedDate":"2021-12-07T20:52:00Z","Layer":{},"PkgName":"lodash","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2021-23337","PublishedDate":"2021-02-15T13:15:00Z","References":["https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-23337","https://github.com/advisories/GHSA-35jh-r3h4-6jhm","https://github.com/lodash/lodash/blob/ddfd9b11a0126db2302cb70ec9973b66baec0975/lodash.js#L14851","https://github.com/lodash/lodash/blob/ddfd9b11a0126db2302cb70ec9973b66baec0975/lodash.js%23L14851","https://github.com/lodash/lodash/commit/3469357cff396a26c363f8c1b5a91dde28ba4b1c","https://nvd.nist.gov/vuln/detail/CVE-2021-23337","https://security.netapp.com/advisory/ntap-20210312-0006/","https://snyk.io/vuln/SNYK-JAVA-ORGFUJIONWEBJARS-1074932","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARS-1074930","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARSBOWER-1074928","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARSBOWERGITHUBLODASH-1074931","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARSNPM-1074929","https://snyk.io/vuln/SNYK-JS-LODASH-1040724","https://www.oracle.com//security-alerts/cpujul2021.html","https://www.oracle.com/security-alerts/cpuoct2021.html"],"Severity":"HIGH","SeveritySource":"nvd","Title":"nodejs-lodash: command injection via template","VulnerabilityID":"CVE-2021-23337"},{"CVSS":{"nvd":{"V2Score":4,"V2Vector":"AV:N/AC:L/Au:S/C:N/I:N/A:P","V3Score":6.5,"V3Vector":"CVSS:3.1/AV:N/AC:L/PR:L/UI:N/S:U/C:N/I:N/A:H"},"redhat":{"V3Score":4.4,"V3Vector":"CVSS:3.0/AV:N/AC:H/PR:H/UI:N/S:U/C:N/I:N/A:H"}},"CweIDs":["CWE-770"],"Description":"lodash prior to 4.17.11 is affected by: CWE-400: Uncontrolled Resource Consumption. The impact is: Denial of service. The component is: Date handler. The attack vector is: Attacker provides very long strings, which the library attempts to match using a regular expression. The fixed version is: 4.17.11.","FixedVersion":"4.17.11","InstalledVersion":"4.17.4","LastModifiedDate":"2020-09-30T13:40:00Z","Layer":{},"PkgName":"lodash","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2019-1010266","PublishedDate":"2019-07-17T21:15:00Z","References":["https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2019-1010266","https://github.com/advisories/GHSA-x5rq-j2xg-h7qm","https://github.com/lodash/lodash/commit/5c08f18d365b64063bfbfa686cbb97cdd6267347","https://github.com/lodash/lodash/issues/3359","https://github.com/lodash/lodash/wiki/Changelog","https://nvd.nist.gov/vuln/detail/CVE-2019-1010266","https://security.netapp.com/advisory/ntap-20190919-0004/","https://snyk.io/vuln/SNYK-JS-LODASH-73639"],"Severity":"MEDIUM","SeveritySource":"nvd","Title":"lodash: uncontrolled resource consumption in Data handler causing denial of service","VulnerabilityID":"CVE-2019-1010266"},{"CVSS":{"nvd":{"V2Score":5,"V2Vector":"AV:N/AC:L/Au:N/C:N/I:N/A:P","V3Score":5.3,"V3Vector":"CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:L"},"redhat":{"V3Score":5.3,"V3Vector":"CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:L"}},"Description":"Lodash versions prior to 4.17.21 are vulnerable to Regular Expression Denial of Service (ReDoS) via the toNumber, trim and trimEnd functions.","FixedVersion":"4.17.21","InstalledVersion":"4.17.4","LastModifiedDate":"2021-12-10T18:04:00Z","Layer":{},"PkgName":"lodash","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2020-28500","PublishedDate":"2021-02-15T11:15:00Z","References":["https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-28500","https://github.com/advisories/GHSA-29mw-wpgm-hmr9","https://github.com/lodash/lodash/blob/npm/trimEnd.js#L8","https://github.com/lodash/lodash/blob/npm/trimEnd.js%23L8","https://github.com/lodash/lodash/pull/5065","https://github.com/lodash/lodash/pull/5065/commits/02906b8191d3c100c193fe6f7b27d1c40f200bb7","https://nvd.nist.gov/vuln/detail/CVE-2020-28500","https://security.netapp.com/advisory/ntap-20210312-0006/","https://snyk.io/vuln/SNYK-JAVA-ORGFUJIONWEBJARS-1074896","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARS-1074894","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARSBOWER-1074892","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARSBOWERGITHUBLODASH-1074895","https://snyk.io/vuln/SNYK-JAVA-ORGWEBJARSNPM-1074893","https://snyk.io/vuln/SNYK-JS-LODASH-1018905","https://www.oracle.com//security-alerts/cpujul2021.html","https://www.oracle.com/security-alerts/cpuoct2021.html"],"Severity":"MEDIUM","SeveritySource":"nvd","Title":"nodejs-lodash: ReDoS via the toNumber, trim and trimEnd functions","VulnerabilityID":"CVE-2020-28500"},{"CVSS":{"nvd":{"V2Score":4,"V2Vector":"AV:N/AC:L/Au:S/C:N/I:P/A:N","V3Score":6.5,"V3Vector":"CVSS:3.0/AV:N/AC:L/PR:L/UI:N/S:U/C:N/I:H/A:N"},"redhat":{"V3Score":2.9,"V3Vector":"CVSS:3.0/AV:L/AC:H/PR:N/UI:N/S:U/C:N/I:N/A:L"}},"Description":"lodash node module before 4.17.5 suffers from a Modification of Assumed-Immutable Data (MAID) vulnerability via defaultsDeep, merge, and mergeWith functions, which allows a malicious user to modify the prototype of \"Object\" via __proto__, causing the addition or modification of an existing property that will exist on all objects.","FixedVersion":"4.17.5","InstalledVersion":"4.17.4","LastModifiedDate":"2019-10-03T00:03:00Z","Layer":{},"PkgName":"lodash","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2018-3721","PublishedDate":"2018-06-07T02:29:00Z","References":["https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-3721","https://github.com/advisories/GHSA-fvqr-27wr-82fm","https://github.com/lodash/lodash/commit/d8e069cc3410082e44eb18fcf8e7f3d08ebe1d4a","https://hackerone.com/reports/310443","https://nvd.nist.gov/vuln/detail/CVE-2018-3721","https://security.netapp.com/advisory/ntap-20190919-0004/","https://snyk.io/vuln/npm:lodash:20180130","https://www.npmjs.com/advisories/577"],"Severity":"LOW","SeveritySource":"nodejs-security-wg","Title":"lodash: Prototype pollution in utilities function","VulnerabilityID":"CVE-2018-3721"},{"CVSS":{"nvd":{"V2Score":6.8,"V2Vector":"AV:N/AC:M/Au:N/C:P/I:P/A:P","V3Score":5.6,"V3Vector":"CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L"},"redhat":{"V3Score":5.6,"V3Vector":"CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L"}},"CweIDs":["CWE-20"],"Description":"minimist before 1.2.2 could be tricked into adding or modifying properties of Object.prototype using a \"constructor\" or \"__proto__\" payload.","FixedVersion":"1.2.3, 0.2.1","InstalledVersion":"0.0.8","LastModifiedDate":"2021-07-21T11:39:00Z","Layer":{},"PkgName":"minimist","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2020-7598","PublishedDate":"2020-03-11T23:15:00Z","References":["http://lists.opensuse.org/opensuse-security-announce/2020-06/msg00024.html","https://github.com/advisories/GHSA-vh95-rmgr-6w4m","https://github.com/substack/minimist/commit/38a4d1caead72ef99e824bb420a2528eec03d9ab","https://github.com/substack/minimist/commit/4cf1354839cb972e38496d35e12f806eea92c11f#diff-a1e0ee62c91705696ddb71aa30ad4f95","https://github.com/substack/minimist/commit/63e7ed05aa4b1889ec2f3b196426db4500cbda94","https://linux.oracle.com/cve/CVE-2020-7598.html","https://linux.oracle.com/errata/ELSA-2020-2852.html","https://nvd.nist.gov/vuln/detail/CVE-2020-7598","https://snyk.io/vuln/SNYK-JS-MINIMIST-559764","https://www.npmjs.com/advisories/1179"],"Severity":"MEDIUM","SeveritySource":"nvd","Title":"nodejs-minimist: prototype pollution allows adding or modifying properties of Object.prototype using a constructor or __proto__ payload","VulnerabilityID":"CVE-2020-7598"},{"CVSS":{"nvd":{"V2Score":6.8,"V2Vector":"AV:N/AC:M/Au:N/C:P/I:P/A:P","V3Score":5.6,"V3Vector":"CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L"},"redhat":{"V3Score":5.6,"V3Vector":"CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L"}},"CweIDs":["CWE-20"],"Description":"minimist before 1.2.2 could be tricked into adding or modifying properties of Object.prototype using a \"constructor\" or \"__proto__\" payload.","FixedVersion":"1.2.3, 0.2.1","InstalledVersion":"1.2.0","LastModifiedDate":"2021-07-21T11:39:00Z","Layer":{},"PkgName":"minimist","PrimaryURL":"https://avd.aquasec.com/nvd/cve-2020-7598","PublishedDate":"2020-03-11T23:15:00Z","References":["http://lists.opensuse.org/opensuse-security-announce/2020-06/msg00024.html","https://github.com/advisories/GHSA-vh95-rmgr-6w4m","https://github.com/substack/minimist/commit/38a4d1caead72ef99e824bb420a2528eec03d9ab","https://github.com/substack/minimist/commit/4cf1354839cb972e38496d35e12f806eea92c11f#diff-a1e0ee62c91705696ddb71aa30ad4f95","https://github.com/substack/minimist/commit/63e7ed05aa4b1889ec2f3b196426db4500cbda94","https://linux.oracle.com/cve/CVE-2020-7598.html","https://linux.oracle.com/errata/ELSA-2020-2852.html","https://nvd.nist.gov/vuln/detail/CVE-2020-7598","https://snyk.io/vuln/SNYK-JS-MINIMIST-559764","https://www.npmjs.com/advisories/1179"],"Severity":"MEDIUM","SeveritySource":"nvd","Title":"nodejs-minimist: prototype pollution allows adding or modifying properties of Object.prototype using a constructor or __proto__ payload","VulnerabilityID":"CVE-2020-7598"}]}`
	website223.Status = false
	err = db.Create(&website223).Error
	log.Println(err)

	info43 := database.ScanResult{}
	info43.UUID = "testkrjfnrjnfjrnnf124"
	info43.Method = "HTTP Security Headers"
	info43.CommandName = "oehfoejfjeoi3667124"
	info43.Result = `{
		"Vulnerabilities":{
			"Severity":"else"
		}
	}`
	err = db.Create(&info43).Error
	log.Println(err)

	infoscan := database.RequestInfo{}
	infoscan.ID = 909
	infoscan.UUID = "dkcecckww;kkplwlw0"
	infoscan.RepoLang = ""

	err = db.Create(&infoscan).Error
	log.Println(err)

	websitescan := database.ScanResult{}
	websitescan.ID = 98798
	websitescan.Method = "Website Security"
	websitescan.UUID = "dkcecckww;kkplwlw0"
	websitescan.CommandName = "ororogkijji"
	err = db.Create(&websitescan).Error
	log.Println(err)

}

func TestIntegration(t *testing.T) {
	Integration("test@xenonstack.com", "test", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQyMDYxOTk2LCJpZCI6MjY2LCJuYW1lIjoicmt4c3JhaHVsIiwib3JpZ19pYXQiOjE2NDIwNjAxOTYsInN5c19yb2xlIjoidXNlciJ9.65N1gh51oviJSv5SocebwovJrSMSIO4HQKHrgFwNfg4")
	Integration("test@xenonstack.com", "test", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQyMDYxOTk2LCJpZCI6MjY2LCJuYW1lIjoicmt4c3JhaHVsIiwib3JpZ19pYXQiOjE2NDIwNjAxOTYsInN5c19yb2xlIjoidXNlciJ9.65N1gh51oviJSv5SocebwovJrSMSIO4HQKHrgFwNfg4")

}

func TestWebsiteResult(t *testing.T) {
	info := database.RequestInfo{}
	info.UUID = "testkrjfnrjnfjrnnf1"
	websiteResult(info)
}

func TestGitResult(t *testing.T) {
	data := database.RequestInfo{}
	data.UUID = "testkrjfnrjnfjrnnf12"
	gitResult(data)
	data.UUID = "ig3wih3ii3nhi3iioi3"
	gitResult(data)
	data.UUID = "testkrjfnrjnfjrnnf124"
	gitResult(data)
}

func TestScanInformation(t *testing.T) {
	ScanInformation("dwoidjfdefjejwfhhohjhoi", "test", "test@xenonstack.com", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQyMDYxOTk2LCJpZCI6MjY2LCJuYW1lIjoicmt4c3JhaHVsIiwib3JpZ19pYXQiOjE2NDIwNjAxOTYsInN5c19yb2xlIjoidXNlciJ9.65N1gh51oviJSv5SocebwovJrSMSIO4HQKHrgFwNfg4")
	ScanInformation("dkcecckww;kkplwlw0", "test", "test@xenonstack.com", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQyMDYxOTk2LCJpZCI6MjY2LCJuYW1lIjoicmt4c3JhaHVsIiwib3JpZ19pYXQiOjE2NDIwNjAxOTYsInN5c19yb2xlIjoidXNlciJ9.65N1gh51oviJSv5SocebwovJrSMSIO4HQKHrgFwNfg4")
	ScanInformation("ig3wih3ii3nhi3iioi", "test", "test@xenonstack.com", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQyMDYxOTk2LCJpZCI6MjY2LCJuYW1lIjoicmt4c3JhaHVsIiwib3JpZ19pYXQiOjE2NDIwNjAxOTYsInN5c19yb2xlIjoidXNlciJ9.65N1gh51oviJSv5SocebwovJrSMSIO4HQKHrgFwNfg4")
	ScanInformation("ig3wih3ii33333333", "test", "test@xenonstack.com", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQyMDYxOTk2LCJpZCI6MjY2LCJuYW1lIjoicmt4c3JhaHVsIiwib3JpZ19pYXQiOjE2NDIwNjAxOTYsInN5c19yb2xlIjoidXNlciJ9.65N1gh51oviJSv5SocebwovJrSMSIO4HQKHrgFwNfg4")
	ScanInformation("3", "test", "test@xenonstack.com", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQyMDYxOTk2LCJpZCI6MjY2LCJuYW1lIjoicmt4c3JhaHVsIiwib3JpZ19pYXQiOjE2NDIwNjAxOTYsInN5c19yb2xlIjoidXNlciJ9.65N1gh51oviJSv5SocebwovJrSMSIO4HQKHrgFwNfg4")
	ScanInformation("3", "test", "test@xenonstack.com", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJhaHVsQHhlbm9uc3RhY2suY29tIiwiZXhwIjoxNjQyMDYxOTk2LCJpZCI6MjY2LCJuYW1lIjoicmt4c3JhaHVsIiwib3JpZ19pYXQiOjE2NDIwNjAxOTYsInN5c19yb2xlIjoidXNlciJ9.65N1gh51oviJSv5SocebwovJrSMSIO4HQKHrgFwNfg4")
}
