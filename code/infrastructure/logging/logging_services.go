package logging

import (
	"github.com/OntoLedgy/logging_services/standard_global_logger"
	"log"
)

// TODO replace with interface

var GlobalLogger *log.Logger

func Start_logging(log_folder_name,
	log_file_name_prefix string) {

	var global_logging_service = standard_global_logger.
		Start_logger(log_folder_name,
			log_file_name_prefix)

	GlobalLogger = global_logging_service.Global_logger

}
