package ai.verta.modeldb.versioning.blob;

import ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model.Blob;
import ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model.CodeBlob;
import ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model.ConfigBlob;
import ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model.DatasetBlob;
import ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model.DockerEnvironmentBlob;
import ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model.EnvironmentBlob;
import ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model.EnvironmentVariablesBlob;
import ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model.GitCodeBlob;
import ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model.NotebookCodeBlob;
import ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model.PathDatasetBlob;
import ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model.PythonEnvironmentBlob;
import ai.verta.modeldb.versioning.autogenerated._public.modeldb.versioning.model.S3DatasetBlob;
import java.util.List;
import java.util.function.Function;
import java.util.stream.Collectors;

public class GenerateBlob {

  public static Blob fromProto(ai.verta.modeldb.versioning.Blob blob) {
    if (blob == null) {
      return null;
    }

    Blob obj = new Blob();
    switch (blob.getContentCase()) {
      case DATASET:
        Function<ai.verta.modeldb.versioning.Blob, DatasetBlob> f =
            x -> fromProto(blob.getDataset());
        obj.Dataset = f.apply(blob);
        break;
      case ENVIRONMENT:
        Function<ai.verta.modeldb.versioning.Blob, EnvironmentBlob> fe =
            x -> fromProto(blob.getEnvironment());
        obj.Environment = fe.apply(blob);
        break;
      case CODE:
        Function<ai.verta.modeldb.versioning.Blob, CodeBlob> fc = x -> fromProto(blob.getCode());
        obj.Code = fc.apply(blob);
        break;
      case CONFIG:
        Function<ai.verta.modeldb.versioning.Blob, ConfigBlob> fconf =
            x -> ConfigBlob.fromProto(blob.getConfig());
        obj.Config = fconf.apply(blob);
        break;
    }
    return obj;
  }

  public static DatasetBlob fromProto(ai.verta.modeldb.versioning.DatasetBlob blob) {
    if (blob == null) {
      return null;
    }

    DatasetBlob obj = new DatasetBlob();
    switch (blob.getContentCase()) {
      case S3:
        {
          Function<ai.verta.modeldb.versioning.DatasetBlob, S3DatasetBlob> f =
              x -> S3DatasetBlob.fromProto(blob.getS3());
          obj.S3 = f.apply(blob);
        }
        break;
      case PATH:
        {
          Function<ai.verta.modeldb.versioning.DatasetBlob, PathDatasetBlob> f =
              x -> PathDatasetBlob.fromProto(blob.getPath());
          obj.Path = f.apply(blob);
        }
        break;
    }
    return obj;
  }

  public static EnvironmentBlob fromProto(ai.verta.modeldb.versioning.EnvironmentBlob blob) {
    if (blob == null) {
      return null;
    }

    EnvironmentBlob obj = new EnvironmentBlob();
    switch (blob.getContentCase()) {
      case PYTHON:
        Function<ai.verta.modeldb.versioning.EnvironmentBlob, PythonEnvironmentBlob> f =
            x -> PythonEnvironmentBlob.fromProto(blob.getPython());
        obj.Python = f.apply(blob);
        break;
      case DOCKER:
        Function<ai.verta.modeldb.versioning.EnvironmentBlob, DockerEnvironmentBlob> fd =
            x -> DockerEnvironmentBlob.fromProto(blob.getDocker());
        obj.Docker = fd.apply(blob);
        break;
    }
    {
      Function<ai.verta.modeldb.versioning.EnvironmentBlob, List<EnvironmentVariablesBlob>> f =
          x ->
              blob.getEnvironmentVariablesList().stream()
                  .map(EnvironmentVariablesBlob::fromProto)
                  .collect(Collectors.toList());
      obj.EnvironmentVariables = f.apply(blob);
    }
    {
      Function<ai.verta.modeldb.versioning.EnvironmentBlob, List<String>> f =
          x -> blob.getCommandLineList();
      obj.CommandLine = f.apply(blob);
    }
    return obj;
  }

  public static CodeBlob fromProto(ai.verta.modeldb.versioning.CodeBlob blob) {
    if (blob == null) {
      return null;
    }

    CodeBlob obj = new CodeBlob();
    switch (blob.getContentCase()) {
      case GIT:
        {
          Function<ai.verta.modeldb.versioning.CodeBlob, GitCodeBlob> f =
              x -> GitCodeBlob.fromProto(blob.getGit());
          obj.Git = f.apply(blob);
        }
        break;
      case NOTEBOOK:
        {
          Function<ai.verta.modeldb.versioning.CodeBlob, NotebookCodeBlob> f =
              x -> NotebookCodeBlob.fromProto(blob.getNotebook());
          obj.Notebook = f.apply(blob);
        }
        break;
    }
    return obj;
  }
}