import { DeploymentStatus, type Deployment } from "./types";

export const mockDeployments: Deployment[] = [
    {
        id: '1',
        name: 'nginx-deployment',
        status: DeploymentStatus.Running,
        namespace: 'default',
    },
    {
        id: '2',
        name: 'nginx-deployment',
        status: DeploymentStatus.Running,
        namespace: 'kube-system',
    },
]
