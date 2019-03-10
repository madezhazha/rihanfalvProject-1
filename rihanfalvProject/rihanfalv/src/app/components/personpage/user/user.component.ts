import { Component, OnInit } from '@angular/core';
// import { FlashMessagesService } from 'angular2-flash-messages';
import {GetdataService} from '../../../services/getdata.service';
import {Router} from '@angular/router';
import { DatePipe } from '@angular/common';
import { DomSanitizer } from '@angular/platform-browser';
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
    '../../../../assets/images/1.jpg',
    '../../../../assets/images/2.jpg',
    '../../../../assets/images/3.jpg',
    '../../../../assets/images/4.jpg',
    '../../../../assets/images/5.jpg',
    '../../../../assets/images/6.jpg'
  ];
  msgs: string[] = [
    '信息不完整请补全信息！',
    '两次密码不一致，请重新输入！',
    '信息修改成功!'
  ];
  msg: any ;
  imgsrc: any;
  Password: any;
  isshow: boolean = false;
  issuccess: boolean = false;
  iswarn: boolean = false;
  str: string;
  base64:string;
  hidden: boolean = false; // 隐藏画布
  constructor(
    public router: Router,
    private serve: GetdataService,
    private sanitizer: DomSanitizer,
    // public flashMessagesService: FlashMessagesService,
    private datePipe: DatePipe
    ) { }

  //  跳转到收藏夹
  collection() {
    this.router.navigate(['/collection']);
  }
  //上传本地头像
  uploadImg() {
    const canvas = <HTMLCanvasElement>document.getElementById('canvas');  // 获取canvas标签
    const ctx = canvas.getContext('2d');  // 2D对象，绘图环境
    const uploadimg = <HTMLInputElement>document.getElementById('img');
    const imgfile = uploadimg.files[0];
    const img = new Image();
    let temp = '';
    img.src = window.URL.createObjectURL(imgfile);
    img.onload = () => {
      ctx.drawImage(img, 0, 0, 50, 50);
      this.base64 = canvas.toDataURL('image/png');
      this.imgsrc = this.base64;
      temp = this.base64.substring(22, this.base64.length);
      // ctx.drawImage(img, 100, 100, 200, 200, 0, 0, 100, 100);
      this.temp.Image = temp;
    };
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
          this.serve.change(this.temp).subscribe(() => {
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
    this.serve.get().subscribe(user => {
      this.user = user;
      // 判断传来的是系统头像的路径还是base64
      if(this.user.Image.length > 100){
      let temp: any;
      temp = 'data:image/png;base64, ' + this.user.Image; //给base64添加头缀
      this.user.Image = this.sanitizer.bypassSecurityTrustUrl(temp);
      }
      //存储用户不可修改的信息
      this.temp.UserId = this.user.UserId;
      this.temp.Integral = this.user.Integral;
      this.temp.RegisterDate = this.user.RegisterDate;
      // 转化日期格式
      this.user.RegisterDate = this.datePipe.transform(this.user.RegisterDate, 'yyyy-MM-dd');
      
  });
  }
}
