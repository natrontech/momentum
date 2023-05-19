// create new type Deployment with the following properties
// name: string
// status: string
// namespace: string

export enum DeploymentStatus {
    Running = 'Running',
    Pending = 'Pending',
    Failed = 'Failed',
    Unknown = 'Unknown',
}

export type Deployment = {
    id: string;
    name: string;
    status: DeploymentStatus;
    namespace: string;
}
