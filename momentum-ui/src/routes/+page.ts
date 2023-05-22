import { env } from "$env/dynamic/public";

export function load() {
    // return env.TEST_ENV
    return {
        env: env.PUBLIC_TEST_ENV
    };
}
