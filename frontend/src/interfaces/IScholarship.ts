import { ScholarshipStatusInterface } from "./IScholarshipstatus";
import { ScholarshipTypesInterface } from "./IScholarshiptypes";
export interface ScholarshipInterface {
    ID?: number;
    Name?: string   | null;
    
    StatusID?: number;
    Scholarships_status?: ScholarshipStatusInterface ;
   
    TypeID?: number;
    Scholarships_type?: ScholarshipTypesInterface;
    
    Details:   String  | null;

}