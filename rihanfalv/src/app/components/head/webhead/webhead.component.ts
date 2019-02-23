import { Component, OnInit, } from '@angular/core';
import {InputData} from '../langing/land/input'
import { Input } from '@angular/compiler/src/core';
@Component({
  selector: 'app-webhead',
  templateUrl: './webhead.component.html',
  styleUrls: ['./webhead.component.css']
})
export class WebheadComponent implements OnInit {
  JapanOrKoreaBool:boolean=true;//true代表日本，false代表韩国
  JapanOrKorea:string="日"

  BackgroundImage:string="../../../../assets/背景图片1.png";
  IfLogin:boolean=false;  //是否已经登录


  In:InputData={IfLogin:false,Tip:"",Image};

  constructor() { }

  ngOnInit() {
  }
  //日韩转换的按钮
  JapanKoreaChange(){
    this.JapanOrKoreaBool=!this.JapanOrKoreaBool;
    if(this.JapanOrKoreaBool){
      this.JapanOrKorea="日";
      this.BackgroundImage="../../../../assets/背景图片1.png"
    }
    else{
      this.JapanOrKorea="韩"
      this.BackgroundImage="../../../../assets/背景图片2.png"
    }
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
