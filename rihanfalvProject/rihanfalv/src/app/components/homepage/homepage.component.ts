import { Component, OnInit } from '@angular/core';
import {GolangService} from './golang.service';
import {HotnewsBox, ArticalBox} from './struct';
import * as $ from 'jquery';
var artical_num = 0;  //当前页面上已经有多少文章 
 
@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.css']
})

export class HomepageComponent implements OnInit {
  more_button_show = "显示跟多"; 
  SlideBox_date = HotnewsBox[5];
  artical_date = new Array();   //新闻列表数据
  constructor(
    private service :GolangService
    ) {}
    
    //使用jquery来打到自动轮播的效果，以及调整元素宽度
    ngOnInit() {
    this.get_head_new(); 
    this.get_artical(); 
    var index = 0;  //自动播放框的图片序号，0~4
    var adTimer;  
    setwidth();
    showImg(4);
    $(document).ready(function(){
        //鼠标停留在小方块时显示相应的图片
        //$(".scalle_box dt").mouseover(function() {
       //     index = $(".scalle_box dt").index(this); 
       //     showImg(index);
       //  }).eq(0).mouseover();

        //手机屏幕上左右滑动图片框时图片向左右切换
      //  $(".picturebox").on("swipeleft",function(){
      //       if(index<4) showImg(++index);
      //       restart();
      //   });
      //   $(".picturebox").on("swiperight",function(){
      //       if(index>0) showImg(--index); 
      //       restart();
      //   });

      //滑入停止动画，滑出开始动画.
        $('.picturebox').on("tap",restart).trigger("taphold"); 

        run();
   })//<--ready()函数结束
    function setwidth(){  
      $(".picture, .describe_box").width($(".homepagebody").width() * 5 +"px");
      $(".picture img, .describe").width($(".homepagebody").width()+"px");
    }

    function restart(){
        clearInterval(adTimer);
        run();
    } 
      
    function run(){
        adTimer = setInterval(function() {
          if (++index > 4) {       //最后一张图片之后，转到第一张
              index = 0;
          }
          showImg(index)
        }, 3000);
    }
    function showImg(index) {
        setwidth();
        var imgwidth = $(".picture img").width();
        $(".picture, .describe_box").stop(true, false).animate({
          "margin-left": -imgwidth * index + "px"    //改变 marginTop 属性的值达到轮播的效果
        }, 1000);
        $(".scalle_box dt").removeClass("on")
        .eq(index).addClass("on");
    }
  } //ngoninti()结束


  //加载首页文章列表的数据，每次最多10篇，总共最多100篇，初始化和点击下端加载跟多按钮调用
  get_artical(){
      var temp_array : ArticalBox[];
      if( artical_num > 100 ){                          //主页最多显示的文章数量
        return;    
      }                
      this.service.GetArtical(artical_num).subscribe(   //10是每次加载的文章数量
        result => {
          if (result == null || result.length<10){
            this.more_button_show = "没有跟多了...";
            artical_num=9999;            //服务器里的文章不足100篇
            $(".more").addClass("nomore");
          } 
          //console.log(result.length);
          artical_num += result.length;
          temp_array = result;
          this.artical_date = this.artical_date.concat(temp_array);
        });
        return;
  }
  //加载首页头部的自动轮播框的图片，链接和标题，固定5个,初始化时自动调用
  get_head_new(){
      this.service.GetHeadNews().subscribe(
        result =>{
          //console.log(result)
          this.SlideBox_date = result;
        }
      )
  }
}
