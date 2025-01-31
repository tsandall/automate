/*
This is the root module for the whole app.
It has a lot of imports. To make it easier to keep track of all the various libraries and
components, please keep them in alphabetical order both in the import lists as well as in
the NgModule decorator metadata.
*/

import { BrowserModule } from '@angular/platform-browser';
import { CookieModule } from 'ngx-cookie';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { environment } from '../environments/environment';

// ngrx/store
import { StoreDevtoolsModule } from '@ngrx/store-devtools';
import { StoreModule } from '@ngrx/store';
import { StoreRouterConnectingModule } from '@ngrx/router-store';
import { NgrxEffectsModule } from './ngrx.effects';
import { ngrxReducers, RouterSerializer, runtimeChecks } from './ngrx.reducers';

// angular material stuff
import {
  BrowserAnimationsModule
} from '@angular/platform-browser/animations';

// Modules
import { ApiTokenModule } from './pages/api-token/api-token.module';
import { AppRoutingModule } from './app-routing.module';
import { ChefComponentsModule } from './components/chef-components.module';
import { ChefPipesModule } from './pipes/chef-pipes.module';
import { ComplianceModule } from './pages/+compliance/compliance.module';
import { ComplianceSharedModule } from './pages/+compliance/shared/shared.module';
import { IntegrationsModule } from './pages/integrations/integrations.module';
import { PolicyModule } from './pages/policy/policy.module';
import { ProjectModule } from './pages/project/project.module';
import { RoleModule } from './pages/roles/roles.module';

// Services
import { AttributesService } from './services/attributes/attributes.service';
import { ChefSessionService } from './services/chef-session/chef-session.service';
import { ConfigService } from './services/config/config.service';
import { EventFeedService } from './services/event-feed/event-feed.service';
import { FeatureFlagsService } from './services/feature-flags/feature-flags.service';
import { HttpClientAuthInterceptor } from './services/http/http-client-auth.interceptor';
import { LocalStorageService } from './services/storage/localstorage.service';
import { MetadataService } from './services/metadata/metadata.service';
import { NodeDetailsResolverService } from './services/node-details/node-details-resolver.service';
import { NodeDetailsService } from './services/node-details/node-details.service';
import {
  NodeNoRunsDetailsResolverService
} from './services/node-details/node-noruns-details-resolver.service';
import { NodeRunsService } from './services/node-details/node-runs.service';
import { ProjectService } from './entities/projects/project.service';
import { ProjectsFilterService } from './services/projects-filter/projects-filter.service';
import { RulesService } from './services/rules/rules.service';
import { RunHistoryStore } from './services/run-history-store/run-history.store';
import { SessionStorageService } from './services/storage/sessionstorage.service';
import { SidebarService } from './services/sidebar/sidebar.service';
import { TelemetryService } from './services/telemetry/telemetry.service';

// Requests
import { ApiTokenRequests } from './entities/api-tokens/api-token.requests';
import { AutomateSettingsRequests } from './entities/automate-settings/automate-settings.requests';
import { ClientRunsRequests } from './entities/client-runs/client-runs.requests';
import { CredentialRequests } from './entities/credentials/credential.requests';
import { JobRequests } from './entities/jobs/job.requests';
import { LicenseStatusRequests } from './entities/license/license.requests';
import { ManagerRequests } from './entities/managers/manager.requests';
import { PolicyRequests } from './entities/policies/policy.requests';
import { ProfileRequests } from './entities/profiles/profile.requests';
import { ProjectRequests } from './entities/projects/project.requests';
import { ServiceGroupsRequests } from './entities/service-groups/service-groups.requests';
import { RoleRequests } from './entities/roles/role.requests';
import { RuleRequests } from './entities/rules/rule.requests';
import { TeamRequests } from './entities/teams/team.requests';
import { UserPermsRequests } from './entities/userperms/userperms.requests';
import { UserRequests } from './entities/users/user.requests';
import { ProjectsFilterRequests } from './services/projects-filter/projects-filter.requests';


// Helpers
import { HistorySelection } from './helpers/history-selection/history-selection';
import { UserDetailsNonAdminResolve } from './pages/user-details/user-details.resolver';

// Page Components
import { AppComponent } from './app.component';
import { ApplicationsComponent } from './pages/applications/applications.component';
import { AttributesComponent } from './page-components/attributes/attributes.component';
import { AutomateSettingsComponent } from './pages/automate-settings/automate-settings.component';
import { ClientRunsComponent } from './pages/client-runs/client-runs.component';
import {
  ClientRunsSidebarComponent
} from './page-components/client-runs-sidebar/client-runs-sidebar.component';
import {
  ClientRunsTableComponent
} from './page-components/client-runs-table/client-runs-table.component';
import {
  ConvergeRadialGraphComponent
} from './page-components/converge-radial-graph/converge-radial-graph.component';
import {
  CreateV1TeamModalComponent
} from './page-components/create-v1-team-modal/create-v1-team-modal.component';
import { DateSelectorComponent } from './page-components/date-selector/date-selector.component';
import {
  DeletableNodeControlComponent
} from './page-components/deletable-node-control/deletable-node-control.component';
import {
  DeleteNotificationDialogComponent
} from './page-components/delete-notification-dialog/delete-notification-dialog.component';
import { DeltaViewerComponent } from './page-components/delta-viewer/delta-viewer.component';
import { EventFeedComponent } from './pages/event-feed/event-feed.component';
import {
  EventFeedGuitarStringsComponent
} from './page-components/event-feed-guitar-strings/event-feed-guitar-strings.component';
import {
  EventFeedSidebarComponent
} from './page-components/event-feed-sidebar/event-feed-sidebar.component';
import {
  EventFeedTableComponent
} from './page-components/event-feed-table/event-feed-table.component';
import { EventIconComponent } from './page-components/event-icon/event-icon.component';
import { FeatureFlagsComponent } from './page-components/feature-flags/feature-flags.component';
import { JobAddComponent } from './pages/job-add/job-add.component';
import { JobEditComponent } from './pages/job-edit/job-edit.component';
import { JobListComponent } from './pages/job-list/job-list.component';
import {
  JobNodesFormComponent
} from './page-components/job-nodes-form/job-nodes-form.component';
import {
  JobProfilesFormComponent
} from './page-components/job-profiles-form/job-profiles-form.component';
import {
  JobScheduleFormComponent
} from './page-components/job-schedule-form/job-schedule-form.component';
import { JsonTreeComponent } from './page-components/json-tree/json-tree.component';
import {
  LicenseApplyComponent
} from './page-components/license-apply/license-apply.component';
import {
  LicenseLockoutComponent
} from './page-components/license-lockout/license-lockout.component';
import { LogsModalComponent } from './page-components/logs-modal/logs-modal.component';
import { MiniTableComponent } from './page-components/mini-table/mini-table.component';
import { NavbarComponent } from './page-components/navbar/navbar.component';
import { NodeDetailsComponent } from './pages/node-details/node-details.component';
import {
  NodeNoRunsDetailsComponent
} from './pages/node-noruns-details/node-noruns-details.component';
import { NodeRollupComponent } from './page-components/node-rollup/node-rollup.component';
import { NotificationFormComponent } from './pages/notification-form/notification-form.component';
import { NotificationsComponent } from './pages/notifications/notifications.component';
import {
  ProjectsFilterComponent
} from './page-components/projects-filter/projects-filter.component';
import {
  ProjectsFilterDropdownComponent
} from './page-components/projects-filter-dropdown/projects-filter-dropdown.component';
import { ProfileComponent } from './page-components/profile/profile.component';
import {
  ProfileSidebarComponent
} from './page-components/profile-sidebar/profile-sidebar.component';
import { ResourceItemComponent } from './page-components/resource-item/resource-item.component';
import { ResourcesComponent } from './page-components/resources/resources.component';
import { RunHistoryComponent } from './page-components/run-history/run-history.component';
import { RunListComponent } from './page-components/run-list/run-list.component';
import {
  RunListRoleHeaderComponent
} from './page-components/run-list-role-header/run-list-role-header.component';
import { RunListTableComponent } from './page-components/run-list-table/run-list-table.component';
import { RunSummaryComponent } from './page-components/run-summary/run-summary.component';
import {
  SearchBarComponent
} from './page-components/search-bar/search-bar.component';
import {
  SearchBarFilterBarComponent
} from './page-components/search-bar-filter-bar/search-bar-filter-bar.component';
import {
  SelectListItemComponent
} from './page-components/select-list-item/select-list-item.component';
import { ServiceGroupsComponent } from './pages/service-groups/service-groups.component';
import { SettingsLandingComponent } from './pages/settings-landing/settings-landing.component';
import {
  SidebarSelectListComponent
} from './page-components/sidebar-select-list/sidebar-select-list.component';
import { SigninComponent } from './pages/signin/signin.component';
import { TeamAddUsersComponent } from './pages/team-add-users/team-add-users.component';
import { TeamDetailsComponent } from './pages/team-details/team-details.component';
import {
  ServicesSidebarComponent
} from './page-components/services-sidebar/services-sidebar.component';
import { TeamManagementComponent } from './pages/team-management/team-management.component';
import {
  TelemetryCheckboxComponent
} from './page-components/telemetry-checkbox/telemetry-checkbox.component';
import { UIComponent } from 'app/ui.component';
import { UserDetailsComponent } from './pages/user-details/user-details.component';
import { UserFormComponent } from './pages/user-management/user-form/user-form.component';
import { UserManagementComponent } from './pages/user-management/user-management.component';
import { UserTableComponent } from './page-components/user-table/user-table.component';
import { WelcomeModalComponent } from './page-components/welcome-modal/welcome-modal.component';

@NgModule({
  declarations: [
    // Page Components
    AppComponent,
    ApplicationsComponent,
    AttributesComponent,
    AutomateSettingsComponent,
    ClientRunsComponent,
    ClientRunsSidebarComponent,
    ClientRunsTableComponent,
    ConvergeRadialGraphComponent,
    CreateV1TeamModalComponent,
    DateSelectorComponent,
    DeletableNodeControlComponent,
    DeleteNotificationDialogComponent,
    DeltaViewerComponent,
    EventFeedComponent,
    EventFeedGuitarStringsComponent,
    EventFeedSidebarComponent,
    EventFeedTableComponent,
    EventIconComponent,
    FeatureFlagsComponent,
    JobAddComponent,
    JobEditComponent,
    JobListComponent,
    JobNodesFormComponent,
    JobProfilesFormComponent,
    JobScheduleFormComponent,
    JsonTreeComponent,
    LicenseApplyComponent,
    LicenseLockoutComponent,
    LogsModalComponent,
    MiniTableComponent,
    NavbarComponent,
    NodeDetailsComponent,
    NodeNoRunsDetailsComponent,
    NodeRollupComponent,
    NotificationFormComponent,
    NotificationsComponent,
    ProjectsFilterComponent,
    ProjectsFilterDropdownComponent,
    ProfileComponent,
    ProfileSidebarComponent,
    ResourceItemComponent,
    ResourcesComponent,
    RunHistoryComponent,
    RunListComponent,
    RunListRoleHeaderComponent,
    RunListTableComponent,
    RunSummaryComponent,
    SearchBarComponent,
    SearchBarFilterBarComponent,
    SelectListItemComponent,
    ServiceGroupsComponent,
    ServicesSidebarComponent,
    SettingsLandingComponent,
    SidebarSelectListComponent,
    SigninComponent,
    TeamAddUsersComponent,
    TeamDetailsComponent,
    TeamManagementComponent,
    TelemetryCheckboxComponent,
    UIComponent,
    UserDetailsComponent,
    UserFormComponent,
    UserTableComponent,
    UserManagementComponent,
    WelcomeModalComponent
  ],
  entryComponents: [
    DeleteNotificationDialogComponent
  ],
  imports: [
    ApiTokenModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    BrowserModule,
    ChefComponentsModule,
    ChefPipesModule,
    ComplianceModule,
    ComplianceSharedModule.forRoot(),
    CookieModule.forRoot(),
    FormsModule,
    HttpClientModule,
    IntegrationsModule,
    NgrxEffectsModule,
    PolicyModule,
    ProjectModule,
    ReactiveFormsModule,
    RoleModule,
    StoreModule.forRoot(ngrxReducers, { runtimeChecks }),
    StoreRouterConnectingModule.forRoot({
      serializer: RouterSerializer
    }),
    !environment.production ? StoreDevtoolsModule.instrument({ maxAge: 25 }) : []
  ],
  providers: [
    ApiTokenRequests,
    AttributesService,
    AutomateSettingsRequests,
    ChefSessionService,
    ConfigService,
    ClientRunsRequests,
    CredentialRequests,
    EventFeedService,
    FeatureFlagsService,
    HistorySelection,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: HttpClientAuthInterceptor,
      multi: true
    },
    JobRequests,
    LicenseStatusRequests,
    LocalStorageService,
    ManagerRequests,
    MetadataService,
    NodeDetailsResolverService,
    NodeNoRunsDetailsResolverService,
    NodeDetailsService,
    NodeRunsService,
    PolicyRequests,
    ProfileRequests,
    ProjectRequests,
    ProjectService,
    ProjectsFilterRequests,
    ProjectsFilterService,
    RoleRequests,
    RuleRequests,
    RulesService,
    RunHistoryStore,
    ServiceGroupsRequests,
    SessionStorageService,
    SidebarService,
    TeamRequests,
    TelemetryService,
    UserDetailsNonAdminResolve,
    UserPermsRequests,
    UserRequests
  ],
  bootstrap: [ AppComponent ],
  schemas: [ CUSTOM_ELEMENTS_SCHEMA ]
})

export class AppModule {}
