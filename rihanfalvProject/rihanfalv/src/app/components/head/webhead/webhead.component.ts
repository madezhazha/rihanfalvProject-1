import { Component, OnInit,Output,EventEmitter,Input } from '@angular/core';
import {InputData} from '../langing/land/input'


@Component({
  selector: 'app-webhead',
  templateUrl: './webhead.component.html',
  styleUrls: ['./webhead.component.css']
})
export class WebheadComponent implements OnInit {
  JapanOrKoreaBool:boolean=true;//true代表日本，false代表韩国
  JapanOrKorea:string="日"

  @Output() isJapan=new EventEmitter<boolean>();
  wasJapan:boolean=true;

  BackgroundImage:string="../../../../assets/背景图片1.png";
  @Input() IfLogin:boolean=false;  //是否已经登录


  In:InputData={ID:'',IfLogin:false,Tip:"",Image,Token:''}

  constructor() { }

  ngOnInit() {
    if(localStorage.getItem("JapanOrKorea")=="日"){
      this.JapanOrKorea="日";
      this.BackgroundImage="../../../../assets/背景图片1.png"
      this.wasJapan=true;
      this.JapanOrKoreaBool=true;
    }
    else{
      this.JapanOrKorea="韩"
      this.BackgroundImage="../../../../assets/背景图片2.png"
      this.wasJapan=false;
      this.JapanOrKoreaBool=false;
    }
    if(localStorage.getItem("id")!=null){
      this.In.ID=localStorage.getItem("id");
      this.In.IfLogin=true;
      this.In.Image=localStorage.getItem("headImage");
    }
  }
  //日韩转换的按钮
  JapanKoreaChange(){
    this.JapanOrKoreaBool=!this.JapanOrKoreaBool;
    if(this.JapanOrKoreaBool){
      this.JapanOrKorea="日";
      this.BackgroundImage="../../../../assets/背景图片1.png"
      this.wasJapan=true;
    }
    else{
      this.JapanOrKorea="韩"
      this.BackgroundImage="../../../../assets/背景图片2.png"
      this.wasJapan=false
    }
    // 弹出日韩状态
    this.isJapan.emit(this.wasJapan);
    //将日韩存储到本地，用于组件初始化 
    localStorage.setItem("JapanOrKorea",this.JapanOrKorea)
  }
  login(){
    this.IfLogin=true;
  }

  close(){
    this.IfLogin=false;
  }

  getLoginData(input:InputData){
    this.In=input;
  }

}
