import { ElMessage } from 'element-plus';
class Message {
  private static instance: Message;
  
  private constructor() {}
  
  public static getInstance(): Message {
    if (!this.instance) {
      this.instance = new Message();
    }
    return this.instance;
  }
  
  public success(message: string): void {
    ElMessage({
      type: 'success',
      message,
    });
  }
  
  public error(message: string): void {
    ElMessage({
      type: 'error',
      message,
    });
  }
  
  public warning(message: string): void {
    ElMessage({
      type: 'warning',
      message,
    });
  }
  
  public info(message: string): void {
    ElMessage({
      type: 'info',
      message,
    });
  }
}

const message = Message.getInstance()
export default message;