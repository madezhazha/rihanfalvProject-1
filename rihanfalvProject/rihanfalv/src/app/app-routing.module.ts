import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { FeedbackComponent } from './components/feedback/feedback.component';
import { FeedbacksuccessComponent } from './components/Feedback/feedbacksuccess/feedbacksuccess.component';
import { PersonalfeedbackComponent } from './components/Feedback/personalfeedback/personalfeedback.component';
import { LegalComponent } from './components/provisions/legal/legal.component';
import { ArticleComponent } from './components/provisions/article/article.component';
import { ContentComponent } from './components/provisions/content/content.component';
import { PaperComponent } from './components/paper/paper.component';
import { PaperwebComponent } from './components/paper/paperweb/paperweb.component';
import { CaseComponent } from './components/case/case.component';
import { CaseDataComponent } from './components/case/case-data/case-data.component';

import { PopularComponent } from '../app/components/chat/popular/popular.component';
import { PostComponent } from '../app/components/chat/post/post.component';
import { ReplyPageComponent } from '../app/components/chat/reply-page/reply-page.component';
import { ChatHeadComponent } from '../app/components/chat/chat-head/chat-head.component';
import { TagComponent } from '../app/components/chat/tag/tag.component';
import { HomepageComponent} from './components/homepage/homepage.component';
import { SearchComponent } from './components/search/search.component';
import { SearchresultComponent } from './components/search/searchresult/searchresult.component';

import { MychatComponent } from './components/chat/mychat/mychat.component';
import { ToQuestionComponent } from './components/chat/to-question/to-question.component';
import { MyQuestionComponent } from './components/chat/my-question/my-question.component';
import { MyAmswerComponent } from './components/chat/my-amswer/my-amswer.component';

import { UserComponent } from './components/personpage/user/user.component' ;
import { CollectionComponent } from './components/personpage/collection/collection.component' ;
import { UploadHeadimageComponent } from './components/personpage/upload-headimage/upload-headimage.component';

const routes: Routes = [
  // { path: 'text', component:PaperComponent },               //测试专用
  { path: 'homepage', component: HomepageComponent},        //首页
  { path: 'discussionarea', component: PopularComponent },  //讨论区
  { path: 'legal', component: LegalComponent  },             //法律条文
  { path: 'caseanalysis', component: CaseComponent},          //案例分析页面
  { path: 'paper', component: PaperComponent },             //相关论文
  { path: 'Feedback', component: FeedbackComponent},
  { path: 'Personalfeedback', component: PersonalfeedbackComponent },
  { path: 'Feedbacksuccess', component: FeedbacksuccessComponent },
  { path: 'article', component: ArticleComponent },
  { path: 'content', component: ContentComponent },
  { path: 'paperweb/:route/:ArticleID', component: PaperwebComponent},
  { path: 'display-data', component: CaseDataComponent},
  { path: 'post', component: PostComponent },
  { path: 'replyPage', component: ReplyPageComponent },
  { path: 'head', component: ChatHeadComponent },
  { path: 'tag', component: TagComponent },
  { path: 'search', component: SearchComponent },
  { path: 'searchresult', component: SearchresultComponent },

  { path: 'mychat', component: MychatComponent },
  { path: 'toquestion', component: ToQuestionComponent },
  { path: 'myquestion', component: MyQuestionComponent },
  { path: 'myanswer', component: MyAmswerComponent },


  {path: 'userpage', component: UserComponent},
  {path: 'collection', component: CollectionComponent},
  {path: 'uploadimage', component: UploadHeadimageComponent},


  { path: '', redirectTo: 'homepage', pathMatch: 'full' } ,
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
