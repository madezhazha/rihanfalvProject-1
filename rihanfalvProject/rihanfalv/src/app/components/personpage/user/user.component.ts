import { Component, OnInit ,ElementRef,ViewChild} from '@angular/core';
// import { FlashMessagesService } from 'angular2-flash-messages';
import { GetdataService } from '../../../services/getdata.service';
import { Router } from '@angular/router';
import { DatePipe } from '@angular/common';
import { DomSanitizer } from '@angular/platform-browser';
import { WebheadComponent } from '../../head/webhead/webhead.component';
@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  @ViewChild(WebheadComponent)
  head:WebheadComponent;
  user: any;
  temp = {
    UserId: 0,
    UserName: '',
    Email: '',
    Password: '',
    Image: '',
    Integral: 0,
    RegisterDate: ''
  };
  imgsrcs: string[] = [
    'assets/images/1.jpg',
    'assets/images/2.jpg',
    'assets/images/3.jpg',
    'assets/images/4.jpg',
    'assets/images/5.jpg',
    'assets/images/6.jpg'
  ];
  msgs: string[] = [
    '信息不完整请补全信息！',
    '两次密码不一致，请重新输入！',
    '信息修改成功!'
  ];
  msg: any;
  imgsrc: any;
  Password: any;
  isshow: boolean = false;
  issuccess: boolean = false;
  iswarn: boolean = false;
  str: string;
  base64: string;
  hidden: boolean = false; // 隐藏画布
  id: number = 0;  //登录后获取id
  money:string='5';   //支付金额
  integral:string='50';   //支付金额对应的积分
  constructor(
    public router: Router,
    private serve: GetdataService,
    private sanitizer: DomSanitizer,
    // public flashMessagesService: FlashMessagesService,
    private datePipe: DatePipe
  ) { }

  sumbitFive(){
    this.money="5";
    this.integral = '50';
  }

  sumbitTen(){
    this.money="10";
    this.integral = '120';
  }

  sumbitTwenty(){
    this.money="20";
    this.integral = '300'
  }

  sumbitFifty(){
    this.money="50";
    this.integral = '800'
  }
  //  跳转到收藏夹
  collection() {
    this.router.navigate(['/collection']);
  }
  //跳转到反馈
  feedback() {

    this.router.navigate(['/Personalfeedback/'])
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
    console.log(img.src);
    img.onload = () => {
      ctx.drawImage(img, 0, 0, 50, 50);
      this.base64 = canvas.toDataURL('image/png');
      console.log(this.base64);
      this.imgsrc = this.base64;
      temp = this.base64.substring(22, this.base64.length);
      // ctx.drawImage(img, 100, 100, 200, 200, 0, 0, 100, 100);
      this.temp.Image = temp;
    };
  }
  // 修改信息
  changeinfo() {
    // console.log('修改');
    if (this.temp.UserName == '' || this.temp.Password == '') {
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
            // 修改完成后调用头部的初始化函数
            this.ngOnInit();
            this.head.ngOnInit();
          });
          if(this.temp.Image.length<50){
            localStorage.setItem('headImage', this.temp.Image);
          }else{
            localStorage.setItem('headImage', 'data:image/jpg;base64,'+this.temp.Image);
          }
          
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
    this.Password = '';
    this.imgsrc = '';
    this.issuccess = false;
    this.iswarn = false;
    this.isshow = false;
  }
  //积分充值
  recharge(){
    this.temp.Integral = this.temp.Integral + Number(this.integral);
    this.user.Integral = this.temp.Integral;
    // 去掉头像base64的前缀
    let tempimg: string;
    tempimg = localStorage.getItem('headImage');
    this.user.Image = tempimg.substring(22,tempimg.length);
    // console.log(tempimg);
    this.serve.change(this.user).subscribe(()=>{
      this.user.Image = tempimg;
    });

  }

  ngOnInit() {
    this.id = JSON.parse(localStorage.getItem('id'));
    console.log(this.id);
    this.serve.get(this.id).subscribe(user => {
      this.user = user;
      console.log(this.user);
      // 判断传来的是系统头像的路径还是base64
      if (this.user.Image.length > 100) {
        let temp: any;
        temp = 'data:image/jpg;base64, ' + this.user.Image; //给base64添加头缀
        this.user.Image = this.sanitizer.bypassSecurityTrustUrl(temp);
      }
      //存储用户不可修改的信息
      this.temp.UserId = this.user.UserId;
      this.temp.Integral = this.user.Integral;
      // console.log(this.temp.Integral);
      this.temp.RegisterDate = this.user.RegisterDate;
      this.temp.Email = this.user.Email;
      // 转化日期格式
      this.user.RegisterDate = this.datePipe.transform(this.user.RegisterDate, 'yyyy-MM-dd');

    });
  }
}
