import { Configuration } from "$lib/momentum-core-client";
import { ApplicationsApi, DeploymentsApi, RepositoriesApi, StagesApi } from "$lib/momentum-core-client/api";
import type { RepositoriesResponse } from "$lib/pocketbase/generated-types"

export function buildBasePath(repository: RepositoriesResponse): string {

    let basePath = "http://"

    basePath += repository.coreHost;
    basePath += repository.corePort ? ":" + repository.corePort : "";
    basePath += repository.coreBasePath.startsWith("/") ? repository.coreBasePath : "/" + repository.coreBasePath;

    return basePath;
}

export function createRepositoryClient(repository: RepositoriesResponse): RepositoriesApi {

    let service = new RepositoriesApi(new Configuration(), buildBasePath(repository));

    return service;
}

export function createApplicationClient(repository: RepositoriesResponse): ApplicationsApi {

    let service = new ApplicationsApi(new Configuration(), buildBasePath(repository));

    return service;
}

export function createStageClient(repository: RepositoriesResponse): StagesApi {

    let service = new StagesApi(new Configuration(), buildBasePath(repository));

    return service;
}

export function createDeploymentClient(repository: RepositoriesResponse): DeploymentsApi {

    let service = new DeploymentsApi(new Configuration(), buildBasePath(repository));

    return service;
}
