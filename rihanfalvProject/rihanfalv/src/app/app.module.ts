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
//import { GolangService } from './components/homepage/golang.service';

import { FeedbacksuccessComponent } from './components/feedback/feedbacksuccess/feedbacksuccess.component';
import { PersonalfeedbackComponent } from './components/feedback/personalfeedback/personalfeedback.component';

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
import { ApiSerivice } from './services/apiservice';



// 个人主页
import { UserComponent } from './components/personpage/user/user.component' ;
import { CollectionComponent } from './components/personpage/collection/collection.component' ;
import { ThesisComponent } from './components/personpage/collection/thesis/thesis.component' ;
import { DatePipe } from '@angular/common';
import { LoginserviceService} from '../app/services/loginservice.service';
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
import { SearchresultComponent } from './components/search/searchresult/searchresult.component';
import { SearchkeywordPipe } from './pipes/searchkeyword.pipe';
//我的问答
import { MychatComponent } from './components/chat/mychat/mychat.component';
import { ToQuestionComponent } from './components/chat/to-question/to-question.component';
import { MyQuestionComponent } from './components/chat/my-question/my-question.component';
import { MyAmswerComponent } from './components/chat/my-amswer/my-amswer.component';

import { CollectedcaseComponent } from './components/personpage/collection/collectedcase/collectedcase.component';
import { CollectedtopicComponent } from './components/personpage/collection/collectedtopic/collectedtopic.component';
import { CasethingComponent } from './components/case/casething/casething.component';

import {HashLocationStrategy , LocationStrategy} from '@angular/common';

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
    TagComponent,
    SearchresultComponent,
    SearchkeywordPipe,
    CollectedcaseComponent,
    CollectedtopicComponent,
    MychatComponent,
    ToQuestionComponent,
    MyQuestionComponent,
    MyAmswerComponent,
    CasethingComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,

    ReactiveFormsModule,
    HttpClientModule,
    FormsModule,
  ],
  providers: [
    {provide:LocationStrategy,useClass:HashLocationStrategy},
    ApiSerivice,
    GetdataService,
    DatePipe,
    LoginserviceService,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
