package tinycmdb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// CreateDB creates a database and sets default and technical values internally used
func CreateDB(basepath string) error {
	db, err := sql.Open("sqlite3", basepath+"tinycmdb.db?_fk=true")
	e(err)
	defer db.Close()

	sqlStmt := `create table tinycmdb (version text);`
	_, err = db.Exec(sqlStmt)
	dbe(err, sqlStmt)

	// The configuration item table where mostly all other tables relate to
	sqlStmt = `create table ci (id integer not null, 
				uuid text unique not null, 
				shortname text not null, 
				longdesc text,
				citype text,
				tschg text,
				);`
	_, err = db.Exec(sqlStmt)
	dbe(err, sqlStmt)

	// Metadata for configuration items
	sqlStmt = `create table cimetadata (id integer not null,
				manufacturer text,
				serialno text,
				productiondate text,
				purchasedate text,
				activefrom text,
				activeto text,
				license text,
				tschg text,
				ciuuid text,
				foreign key(ciuuid) references ci(uuid)
				);`
	_, err = db.Exec(sqlStmt)
	dbe(err, sqlStmt)

	// A generic hardware table
	sqlStmt = `create table hwtable (id integer not null,
				uuid text unique not null, 
				cpucount text,
				cputype text,
				cpucores text,
				memorysize text,
				memorytype text,
				storagesize text,
				storagetype text,
				powerconsumption integer,
				tschg text,
				ciuuid text,
				foreign key(ciuuid) references ci(uuid)
				);`
	_, err = db.Exec(sqlStmt)
	dbe(err, sqlStmt)

	// A generic software table
	sqlStmt = `create table swtable (id integer not null,
				uuid text unique not null, 
				shortname text,
				longname text,
				manufacturer text,
				version text,
				patchlevel text,
				year text,
				platform text,
				tschg text,
				ciuuid text,
				foreign key(ciuuid) references ci(uuid)
				)`
	_, err = db.Exec(sqlStmt)
	dbe(err, sqlStmt)

	// A generic table for network configurations
	sqlStmt = `create table netconf (id integer not null,
				uuid text unique not null, 
				ipaddr text,
				macaddr text,
				subnet text,
				gateway text,
				dns text,
				dhcpd text,
				ciuuid text,
				tschg text,
				netseguuid text,
				foreign key(netseguuid) references netsegment(uuid),
				foreign key(ciuuid) references ci(uuid)
				);`
	_, err = db.Exec(sqlStmt)
	dbe(err, sqlStmt)

	// A generic table for basic documents in text form or as blobs
	sqlStmt = `create table cidoc (id integer not null,
				uuid text unique not null,
				doctype text,
				subject text,
				citxt text,
				scan blob,
				tschg text,
				ciuuid text,
				foreign key(ciuuid) references ci(uuid)
				);`
	_, err = db.Exec(sqlStmt)
	dbe(err, sqlStmt)

	// A table for generic netsegment configurations
	sqlStmt = `create table netsegment (id integer not null,
				netname text,
				gateway text,
				firstaddr text,
				lastaddr text
				tschg text,
				);`
	_, err = db.Exec(sqlStmt)
	dbe(err, sqlStmt)

	// A table for physical and logical locations
	sqlStmt = `create table cilocation (id integer not null,
				uuid text unique not null,
				locname text,
				shortdesc text,
				country text,
				city text,
				street text,
				number text,
				room text,
				row text,
				shelf text,
				rack text,
				misc text,
				uri text,
				tschg text,
				ciuuid text,
				foreign key(ciuuid) references ci(uuid)
				);`
	_, err = db.Exec(sqlStmt)
	dbe(err, sqlStmt)


	// A system table to connect configuration type items
	sqlStmt = `create table systems(id integer not null,
				uuid text unique not null,
				thisuuid text,
				hasuuid text,
				tschg text,
				);`

	// Insert default values
	_, err = db.Exec("insert into tinycmdb(version) values('1.0')")
	dbe(err,"Insert default value")

	_, err = db.Exec("insert into netsegment(id, uuid, netname, gateway, firstaddr, lastaddr) values(0, '0000-0000-0000-0000', 'dummy', null,null, null)")
	dbe(err,"Insert default value")

	return err
}
