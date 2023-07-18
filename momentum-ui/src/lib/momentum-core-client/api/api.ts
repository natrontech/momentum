export * from './applications.service';
import { ApplicationsService } from './applications.service';
export * from './deployments.service';
import { DeploymentsService } from './deployments.service';
export * from './repositories.service';
import { RepositoriesService } from './repositories.service';
export * from './stages.service';
import { StagesService } from './stages.service';
export const APIS = [ApplicationsService, DeploymentsService, RepositoriesService, StagesService];
