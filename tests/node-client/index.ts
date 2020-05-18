import * as grpc from '@grpc/grpc-js';
import * as continersPb from "./grpc/containers_grpc_pb";
import { IContainersClient } from './grpc/containers_grpc_pb';
import { ListRequest, ListResponse } from "./grpc/containers_pb";

const ContainersServiceClient = grpc.makeClientConstructor(continersPb["com.docker.api.containers.v1.Containers"], "ContainersClient");
const client = new ContainersServiceClient("unix:///tmp/backend.sock", grpc.credentials.createInsecure()) as unknown as IContainersClient;

let backend = process.argv[2] || "moby";
const meta = new grpc.Metadata();
meta.set("CONTEXT_KEY", backend);

client.list(new ListRequest(), meta, (err: any, response: ListResponse) => {
    if (err != null) {
        console.error(err);
        return;
    }

    const containers = response.getContainersList();

    containers.forEach(container => {
        console.log(container.getId(), container.getImage());
    });
});