
db_default: ./.sqliterc
	rm -f DCbG-caller/database.db
	@for x in $(shell ls *.sql); do sqlite3 DCbG-caller/database.db < $$x > /dev/null; done

clean:
	rm -f DCbG-caller/database.db

shell:	$(HOME)/.sqliterc
	@rm -f database.db
	@for x in $(shell ls *.sql); do sqlite3 database.db < $$x > /dev/null; done
ifeq (, $(shell which litecli))
	sqlite3 database.db
else
	litecli database.db
endif