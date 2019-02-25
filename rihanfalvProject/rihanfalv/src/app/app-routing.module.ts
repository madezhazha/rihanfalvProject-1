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

const routes: Routes = [
  // { path: 'text', component:PaperComponent },               //测试专用
  { path: 'homepage', component: HomepageComponent},        //首页
  { path: 'discussionarea', component: PopularComponent },  //讨论区
  { path: 'legal', component:LegalComponent  },             //法律条文
  { path: 'caseanalysis',component:CaseComponent},          //案例分析页面
  { path: 'paper', component: PaperComponent },             //相关论文
  { path: 'Feedback', component: FeedbackComponent},
  { path: 'Personalfeedback', component: PersonalfeedbackComponent },
  { path: 'Feedbacksuccess', component: FeedbacksuccessComponent },
  { path: 'legal', component: LegalComponent },
  { path: 'article', component: ArticleComponent },
  { path: 'content', component: ContentComponent },
  { path: 'paperweb/:ID', component: PaperwebComponent},
  { path: 'display-data',component:CaseDataComponent},
  { path: 'post', component: PostComponent },
  { path: 'replyPage', component: ReplyPageComponent },
  { path: 'head', component: ChatHeadComponent },
  { path: 'tag', component: TagComponent },
  { path: '', redirectTo: 'homepage', pathMatch: 'full' } ,
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
