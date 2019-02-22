import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-upload-headimage',
  templateUrl: './upload-headimage.component.html',
  styleUrls: ['./upload-headimage.component.css']
})
export class UploadHeadimageComponent implements OnInit {
  /*********这里是数据区****** */
  reader=new FileReader();//图片读取对象
  canvasImg:string='';//储存导出的图片的base64
  /*********这里是函数区****** */
  constructor() { }

  ngOnInit() {
  }


  change(){
    //var img=<HTMLImageElement>document.getElementById('img')
    var img=new Image();//新建一个图片对象，储存导入的图片
    var image=<HTMLInputElement>document.getElementById('image')
    this.reader.readAsDataURL(image.files[0]);//将图片导入
    var canvas = <HTMLCanvasElement>document.getElementById('tutorial');
    var ctx = canvas.getContext('2d');
    this.reader.onload = ()=>{
      //读取完成后，将结果赋值给img的src
      img.src=<string>this.reader.result;
      var headImageWidth;//头像的宽度；
      headImageWidth=Math.min(1000,img.height,img.width)//存储画布的大小，最大为1000，最小为图片宽度和高度的最小值
      canvas.width=headImageWidth;//画布的高度
      canvas.height=headImageWidth;//画布的高度
      /**在画布描绘图片 */
      ctx.drawImage(img,(img.width-headImageWidth)/2,(img.height-headImageWidth)/2,headImageWidth,headImageWidth,0,0,headImageWidth,headImageWidth);
      /**将图片导出成base64 */
      this.canvasImg=canvas.toDataURL();

    }
  
  }

}
