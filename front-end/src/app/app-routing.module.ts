import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './auth-guard/auth.guard';

const routes: Routes = [
  {
    //following Thomas Oliver tutorial, Lazy Loading public
    path:'public',
    loadChildren: () => import('./public/public.module').then(m => m.PublicModule)
  },

  {
    //same idea as public lazy load but this also includes jwt checking for login/access
    path:'protected',
    canActivate: [AuthGuard],
    loadChildren: () => import('./protected/protected.module').then(m => m.ProtectedModule)
  },

  {
    path:'**',
    redirectTo:'public',
    pathMatch:'full'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
