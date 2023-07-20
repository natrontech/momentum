import { Configuration } from '$lib/momentum-core-client';
import { HttpClient, HttpXhrBackend } from '@angular/common/http';

export class HttpClientFactory {

    public createHttpClient(): HttpClient {
        return new HttpClient(new HttpXhrBackend({ build: () => new XMLHttpRequest() }));
    }

    public createConfiguration(): Configuration {
        return new Configuration()
    }
}