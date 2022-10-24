import { ReportInterface } from "./IReport";
import { StatusInterface } from "./IStatus";

export interface StudentlistInterface {
    ID?: number,
    ReportID? : number,
	  Report?   :  ReportInterface,
    Reason?: string,
    Amount?: number,


	  StatusID? : number,
	  Status?   : StatusInterface,
    SaveTime?  : Date | null,
  }
  