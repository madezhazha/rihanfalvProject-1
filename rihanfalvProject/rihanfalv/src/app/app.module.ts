import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { ReactiveFormsModule } from '@angular/forms';
import { FormsModule } from '@angular/forms';
import { HttpClientModule} from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { FeedbackComponent } from './components/feedback/feedback.component';
import { PaperComponent } from './components/paper/paper.component';
import { PaperwebComponent } from './components/paper/paperweb/paperweb.component';
import { SearchComponent } from './components/search/search.component';
import { CaseComponent } from './components/case/case.component';
import { PersonpageComponent } from './components/personpage/personpage.component';
import { HeadComponent } from './components/head/head.component';
import { HomepageComponent} from './components/homepage/homepage.component';

import { FeedbacksuccessComponent } from './components/Feedback/feedbacksuccess/feedbacksuccess.component';
import { PersonalfeedbackComponent } from './components/Feedback/personalfeedback/personalfeedback.component';

// ng头部组件
// 头部组件
import { LangingComponent } from './components/head/langing/langing.component';
import { LandComponent } from './components/head/langing/land/land.component';
import { RegisterComponent } from './components/head/langing/register/register.component';
import { ForgetpasswordComponent } from './components/head/langing/forgetpassword/forgetpassword.component';
import { WebheadComponent } from './components/head/webhead/webhead.component';
import { UploadHeadimageComponent } from './components/personpage/upload-headimage/upload-headimage.component';

import { LegalComponent } from './components/provisions/legal/legal.component';
import { ArticleComponent } from './components/provisions/article/article.component';
import { ContentComponent } from './components/provisions/content/content.component';
import { ApiSerivice } from './apiservice';


// 个人主页
import { UserComponent } from './components/personpage/user/user.component' ;
import { CollectionComponent } from './components/personpage/collection/collection.component' ;
import { ThesisComponent } from './components/personpage/collection/thesis/thesis.component' ;
import { DatePipe } from '@angular/common';
// 服务
import {GetdataService} from '../app/services/getdata.service';
import { SanitizeHtmlPipe } from './pipes/sanitize-html.pipe';
import { CaseDataComponent } from './components/case/case-data/case-data.component';
//讨论区
import { PostComponent } from './components/chat/post/post.component';
import { PopularComponent } from './components/chat/popular/popular.component';
import { ReplyPageComponent } from './components/chat/reply-page/reply-page.component';
import { ChatHeadComponent } from './components/chat/chat-head/chat-head.component';
import { TagComponent } from './components/chat/tag/tag.component';


@NgModule({
  declarations: [
    AppComponent,
    FeedbackComponent,
    PaperComponent,
    PaperwebComponent,
    SearchComponent,
    CaseComponent,
    PersonpageComponent,
    HeadComponent,
    HomepageComponent,
    // 头部组件
    LangingComponent,
    LandComponent,
    RegisterComponent,
    ForgetpasswordComponent,
    WebheadComponent,
    UserComponent,
    CollectionComponent,
    ThesisComponent,
    FeedbacksuccessComponent,
    PersonalfeedbackComponent,
    UploadHeadimageComponent,
    LegalComponent,
    ArticleComponent,
    ContentComponent,
    SanitizeHtmlPipe,
    CaseDataComponent,
    //讨论区
    PostComponent,
    PopularComponent,
    ReplyPageComponent,
    ChatHeadComponent,
    TagComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,

    ReactiveFormsModule,
    HttpClientModule,
    FormsModule,
  ],
  providers: [
    ApiSerivice,
    GetdataService,
     DatePipe
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
