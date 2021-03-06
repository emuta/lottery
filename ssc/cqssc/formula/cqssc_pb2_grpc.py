# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import cqssc_pb2 as cqssc__pb2


class FormulaStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Validate = channel.unary_unary(
        '/cqssc.Formula/Validate',
        request_serializer=cqssc__pb2.ValidateReq.SerializeToString,
        response_deserializer=cqssc__pb2.ValidateResp.FromString,
        )
    self.Compute = channel.unary_unary(
        '/cqssc.Formula/Compute',
        request_serializer=cqssc__pb2.ComputeReq.SerializeToString,
        response_deserializer=cqssc__pb2.ComputeResp.FromString,
        )


class FormulaServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Validate(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Compute(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_FormulaServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Validate': grpc.unary_unary_rpc_method_handler(
          servicer.Validate,
          request_deserializer=cqssc__pb2.ValidateReq.FromString,
          response_serializer=cqssc__pb2.ValidateResp.SerializeToString,
      ),
      'Compute': grpc.unary_unary_rpc_method_handler(
          servicer.Compute,
          request_deserializer=cqssc__pb2.ComputeReq.FromString,
          response_serializer=cqssc__pb2.ComputeResp.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'cqssc.Formula', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
