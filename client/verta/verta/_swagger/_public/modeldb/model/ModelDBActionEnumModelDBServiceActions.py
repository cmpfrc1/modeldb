# THIS FILE IS AUTO-GENERATED. DO NOT EDIT
from verta._swagger.base_type import BaseType

class ModelDBActionEnumModelDBServiceActions(BaseType):
  _valid_values = [
    "UNKNOWN",
    "ALL",
    "CREATE",
    "READ",
    "UPDATE",
    "DELETE",
    "DEPLOY",
    "PUBLIC_READ",
  ]

  def __init__(self, val):
    if val not in ModelDBActionEnumModelDBServiceActions._valid_values:
      raise ValueError('{} is not a valid value for ModelDBActionEnumModelDBServiceActions'.format(val))
    self.value = val

  def to_json(self):
    return self.value

  def from_json(v):
    if isinstance(v, str):
      return ModelDBActionEnumModelDBServiceActions(v)
    else:
      return ModelDBActionEnumModelDBServiceActions(ModelDBActionEnumModelDBServiceActions._valid_values[v])

