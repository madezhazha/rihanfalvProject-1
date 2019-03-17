import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import { LandService } from './lands/land.service'
import { InputData } from './input';
import { FormControl, FormBuilder } from '@angular/forms';
import { OutputData } from './output'
import { Validators } from '@angular/forms';


@Component({
  selector: 'app-land',
  templateUrl: './land.component.html',
  styleUrls: ['./land.component.css']
})
export class LandComponent implements OnInit {

  @Output() gotoRegisters = new EventEmitter();  //跳转到注册界面
  @Output() gotoForgetPasswords = new EventEmitter(); //跳转到忘记密码界面
  @Output() LoginData = new EventEmitter<InputData>();

  in: InputData = { ID: '', IfLogin: false, Tip: "", Image: '', Token: '' }         //传到子组件（“register”，“forgetPassword”，“land”）
  out: OutputData = { Email: "", Password: "" }             //传到父组件（webhead）的数据
  conversions: any

  profileForm = this.fb.group({
    Email: ['', Validators.required],
    Password: ['', Validators.required]
  })

  constructor(public https: LandService, private fb: FormBuilder) { }

  ngOnInit() {
  }
  gotoRegister() {
    this.gotoRegisters.emit();
  }//“注册”界面跳转

  gotoForgetPassword() {
    this.gotoForgetPasswords.emit();
  }//“忘记密码”界面跳转

  onSubmit() {
    this.out.Email = this.profileForm.value.Email;
    this.out.Password = this.profileForm.value.Password;
    this.https.getInput(this.out).subscribe(
      (data: InputData) => {
        this.in = data;
        if (this.in.Image.length >= 36) {
          this.in.Image = "data:image/jpg;base64," + this.in.Image;
        }
        localStorage.setItem("id", this.in.ID);
        localStorage.setItem('token', this.in.Token);
        localStorage.setItem('headImage', this.in.Image)
        this.LoginData.emit(this.in);
      }
    )
  }//登录功能按钮


}
