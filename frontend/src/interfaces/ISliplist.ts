import { BankingInterface } from "./IBanking";
import { PaymentstatusInterface } from "./IPaymentstatus";
import { StudentListInterface } from "./IStudentlist";

export interface SliplistInterface {
  ID?: number;
  Total: number;
  Slipdate: Date | null;
  BankingID?: number;
  Banking?: BankingInterface;
  PayID?: number;
  Pay?: PaymentstatusInterface;

  StudentListID?: number;
  StudentList?: StudentListInterface;
}