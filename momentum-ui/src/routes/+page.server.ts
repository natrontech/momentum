
import { env } from '$env/dynamic/private';


export function load() {
    // return env.TEST_ENV
    return {
        env: env.TEST_ENV,
    }
}
