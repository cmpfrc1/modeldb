// THIS FILE IS AUTO-GENERATED. DO NOT EDIT
package ai.verta.swagger._public.modeldb.model

import scala.util.Try

import net.liftweb.json._

import ai.verta.swagger._public.modeldb.model.ArtifactTypeEnumArtifactType._
import ai.verta.swagger._public.modeldb.model.OperatorEnumOperator._
import ai.verta.swagger._public.modeldb.model.TernaryEnumTernary._
import ai.verta.swagger._public.modeldb.model.ValueTypeEnumValueType._
import ai.verta.swagger._public.modeldb.model.ProtobufNullValue._
import ai.verta.swagger.client.objects._

case class ModeldbGetExperimentRunCodeVersionResponse (
  code_version: Option[ModeldbCodeVersion] = None
) extends BaseSwagger {
  def toJson(): JValue = ModeldbGetExperimentRunCodeVersionResponse.toJson(this)
}

object ModeldbGetExperimentRunCodeVersionResponse {
  def toJson(obj: ModeldbGetExperimentRunCodeVersionResponse): JObject = {
    new JObject(
      List[Option[JField]](
        obj.code_version.map(x => JField("code_version", ((x: ModeldbCodeVersion) => ModeldbCodeVersion.toJson(x))(x)))
      ).flatMap(x => x match {
        case Some(y) => List(y)
        case None => Nil
      })
    )
  }

  def fromJson(value: JValue): ModeldbGetExperimentRunCodeVersionResponse =
    value match {
      case JObject(fields) => {
        val fieldsMap = fields.map(f => (f.name, f.value)).toMap
        ModeldbGetExperimentRunCodeVersionResponse(
          // TODO: handle required
          code_version = fieldsMap.get("code_version").map(ModeldbCodeVersion.fromJson)
        )
      }
      case _ => throw new IllegalArgumentException(s"unknown type ${value.getClass.toString}")
    }
}
