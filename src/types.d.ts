export interface emailAddress{
    Name: string,
    Address : string
}
  
export interface emailType {
    timestamp: string,
    Body:string,
    Cc: string,
    Content_Transfer_Encoding: string,
    Content_Type: string,
    Date:string,
    ID:number,
    Message_ID: string,
    Mime_Version: string,
    X_FileName: string,
    X_Folder: string,
    X_From: string,
    X_Origin: string,
    X_To: string,
    X_cc: string,
    X_bcc: string,
    from: emailAddress,
    subject: string,
    to: Array<emailAddress>,
}

export interface resultType {
    _index: string, 
    _type: string, 
    _id: string, 
    _score: number, 
    _timestamp: string,
    _source : emailType
  }