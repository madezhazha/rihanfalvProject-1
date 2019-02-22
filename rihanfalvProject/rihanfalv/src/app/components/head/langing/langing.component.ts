import { Component, OnInit ,EventEmitter,Output,Input,OnChanges} from '@angular/core';
import {InputData} from './land/input'
@Component({
  selector: 'app-langing',
  templateUrl: './langing.component.html',
  styleUrls: ['./langing.component.css']
})
export class LangingComponent implements OnInit {
 
  NowFace:number=0;

  @Output() closed= new EventEmitter();    //关闭传值
  @Output() LoginData=new EventEmitter<InputData>();   //传登录数据到webhead（头像等）

  constructor() { }

  ngOnInit() {

  }

  gotoRegister(){
    this.NowFace=1;
  } // 跳转到注册界面

  gotoForgetPassword(){
    this.NowFace=2;
  } //跳转到忘记密码界面

  getLoginData(input:InputData){
    this.LoginData.emit(input);
  }  //从登录组件取得登录的数据

  close(){
    this.closed.emit();
  } 

  getNowFace(data){
    this.NowFace=data;
  }

}
  
