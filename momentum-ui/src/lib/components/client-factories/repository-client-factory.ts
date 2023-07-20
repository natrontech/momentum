import { RepositoriesService } from "$lib/momentum-core-client/api/api"
import type { RepositoriesResponse } from "$lib/pocketbase/generated-types"
import { HttpClientFactory } from "./http-client-factory"

export class RepositoryClientFactory {

    private httpClientFactory = new HttpClientFactory();

    public CreateInstance(repository: RepositoriesResponse): RepositoriesService {

        let httpClient = this.httpClientFactory.createHttpClient();
        let config = this.httpClientFactory.createConfiguration();
        let service = new RepositoriesService(httpClient, this.buildBasePath(repository), config);

        return service;
    }

    public buildBasePath(repository: RepositoriesResponse): string {

        let basePath = "http://"

        basePath += repository.coreHost;
        basePath += repository.corePort ? ":" + repository.corePort : "";
        basePath += repository.coreBasePath.startsWith("/") ? repository.coreBasePath : "/" + repository.coreBasePath;

        return basePath;
    }
}