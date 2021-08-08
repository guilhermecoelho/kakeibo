export class ReportView{
    dateStart: Date;
    dateFinish: Date;
    totalIncome: Float32Array;
    totalExpense: Float32Array;
    balance: Float32Array;
    moviments: Moviment[]
}

export class Moviment{
    date: Date;
    name: string;
    amount: number;
    isExpense:boolean;
}