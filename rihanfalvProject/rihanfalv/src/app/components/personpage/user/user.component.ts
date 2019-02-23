import { Component, OnInit } from '@angular/core';
// import { FlashMessagesService } from 'angular2-flash-messages';
import {GetdataService} from '../../../services/getdata.service';
import {Router} from '@angular/router';
// import { DomSanitizer } from '@angular/platform-browser';
import { DatePipe } from '@angular/common';
@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  user: any;
  temp = {
    UserId: 0,
    UserName: '',
    Email: '',
    Password: '',
    Image: '',
    Integral: '',
    RegisterDate: ''
  };
  imgsrcs: string[] = [
    '../../../assets/image/1.jpg',
    '../../../assets/image/2.jpg',
    '../../../assets/image/3.jpg',
    '../../../assets/image/4.jpg',
    '../../../assets/image/5.jpg',
    '../../../assets/image/6.jpg'
  ];
  msgs: string[] = [
    '信息不完整请补全信息！',
    '两次密码不一致，请重新输入！',
    '信息修改成功!'
  ];
  msg: any ;
  imgsrc: any;
  Password: any;
  isshow: boolean;
  issuccess: boolean;
  iswarn: boolean;
  imageUrl: any;
  str: string;
  base64: '';
  reader = new FileReader();
  canvasimg: '';

  constructor(
    public router: Router,
    private serve: GetdataService,
    // public flashMessagesService: FlashMessagesService,
    // private sanitizer: DomSanitizer,
    private datePipe: DatePipe
    ) {
    this.serve.get().subscribe(user => {
        this.user = user;
        this.temp.UserId = this.user.UserId;
        this.temp.Integral = this.user.Integral;
        this.temp.RegisterDate = this.user.RegisterDate;
        // 转化日期格式
        this.user.RegisterDate = this.datePipe.transform(this.user.RegisterDate, 'yyyy-MM-dd');
    });
  }

  //  跳转到收藏夹
  collection() {
    this.router.navigate(['/collection']);
  }

  // 修改信息
  changeinfo() {
    // console.log('修改');
    if (this.temp.UserName == '' || this.temp.Email == '' || this.temp.Password == '') {
        // alert('请补全信息！');
        this.msg = this.msgs[0];
        this.iswarn = true;
        this.issuccess = false;
    } else {
      if (this.temp.Password != this.Password) {
        // alert('密码不一致，请重新输入！');
        this.msg = this.msgs[1];
        this.iswarn = true;
        this.issuccess = false;
      } else {
        if (this.temp.Image == '') {
          this.temp.Image = this.user.Image;
        } else {
          this.serve.change(this.temp).subscribe(res => {
          });
          // alert('修改成功！');
          this.msg = this.msgs[2];
          this.issuccess = true;
          this.iswarn = false;
        }
      }
    }
  }

  // 选择系统头像，并显示
  changeimg(src) {
    this.imgsrc = src;
    this.temp.Image = src;
  }
  // 关闭修改信息时，清空
  close() {
    this.temp.UserName = '';
    this.temp.Password = '';
    this.temp.Email = '';
    this.Password = '' ;
    this.imgsrc = '';
    this.issuccess = false;
    this.iswarn = false;
    this.isshow = false;
  }
  ngOnInit() {
  }
}
