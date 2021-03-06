# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from transaction.v4 import transaction_service_pb2 as transaction_dot_v4_dot_transaction__service__pb2


class TransactionStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetServiceConfig = channel.unary_unary(
                '/kin.agora.transaction.v4.Transaction/GetServiceConfig',
                request_serializer=transaction_dot_v4_dot_transaction__service__pb2.GetServiceConfigRequest.SerializeToString,
                response_deserializer=transaction_dot_v4_dot_transaction__service__pb2.GetServiceConfigResponse.FromString,
                )
        self.GetMinimumKinVersion = channel.unary_unary(
                '/kin.agora.transaction.v4.Transaction/GetMinimumKinVersion',
                request_serializer=transaction_dot_v4_dot_transaction__service__pb2.GetMinimumKinVersionRequest.SerializeToString,
                response_deserializer=transaction_dot_v4_dot_transaction__service__pb2.GetMinimumKinVersionResponse.FromString,
                )
        self.GetRecentBlockhash = channel.unary_unary(
                '/kin.agora.transaction.v4.Transaction/GetRecentBlockhash',
                request_serializer=transaction_dot_v4_dot_transaction__service__pb2.GetRecentBlockhashRequest.SerializeToString,
                response_deserializer=transaction_dot_v4_dot_transaction__service__pb2.GetRecentBlockhashResponse.FromString,
                )
        self.GetMinimumBalanceForRentExemption = channel.unary_unary(
                '/kin.agora.transaction.v4.Transaction/GetMinimumBalanceForRentExemption',
                request_serializer=transaction_dot_v4_dot_transaction__service__pb2.GetMinimumBalanceForRentExemptionRequest.SerializeToString,
                response_deserializer=transaction_dot_v4_dot_transaction__service__pb2.GetMinimumBalanceForRentExemptionResponse.FromString,
                )
        self.GetHistory = channel.unary_unary(
                '/kin.agora.transaction.v4.Transaction/GetHistory',
                request_serializer=transaction_dot_v4_dot_transaction__service__pb2.GetHistoryRequest.SerializeToString,
                response_deserializer=transaction_dot_v4_dot_transaction__service__pb2.GetHistoryResponse.FromString,
                )
        self.SubmitTransaction = channel.unary_unary(
                '/kin.agora.transaction.v4.Transaction/SubmitTransaction',
                request_serializer=transaction_dot_v4_dot_transaction__service__pb2.SubmitTransactionRequest.SerializeToString,
                response_deserializer=transaction_dot_v4_dot_transaction__service__pb2.SubmitTransactionResponse.FromString,
                )
        self.GetTransaction = channel.unary_unary(
                '/kin.agora.transaction.v4.Transaction/GetTransaction',
                request_serializer=transaction_dot_v4_dot_transaction__service__pb2.GetTransactionRequest.SerializeToString,
                response_deserializer=transaction_dot_v4_dot_transaction__service__pb2.GetTransactionResponse.FromString,
                )


class TransactionServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetServiceConfig(self, request, context):
        """GetServiceConfig returns the service and token parameters for the token.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMinimumKinVersion(self, request, context):
        """GetMinimumKinVersion returns the minimum Kin version that is supported.

        This version will _never_ decrease in non-test scenarios, as it indicates
        a global migration has occured.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetRecentBlockhash(self, request, context):
        """GetRecentBlockhash returns a recent block hash from the underlying network,
        which should be used when crafting transactions. If a transaction fails with
        DuplicateSignature or InvalidNonce, it is recommended that a new block hash
        is retrieved.

        Block hashes are expected to be valid for ~2 minutes.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMinimumBalanceForRentExemption(self, request, context):
        """GetMinimumBalanceForRentExemption returns the minimum amount of lamports that
        must be in an account for it not to be garbage collected.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetHistory(self, request, context):
        """GetHistory returns the transaction history for an account,
        with additional off-chain invoice data, if available.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SubmitTransaction(self, request, context):
        """SubmitTransaction submits a transaction.

        The transaction may include a single Memo[1] instruction.
        If a memo instruction is specified, it must be at position 0
        in the instruction array.

        If an invoice is provided, the Memo instruction must contain a
        Kin Binary memo[2], encoded as base64.

        [1]: https://spl.solana.com/memo
        [2]: https://github.com/kinecosystem/agora-api/blob/master/spec/memo.md
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetTransaction(self, request, context):
        """GetTransaction returns a transaction and additional off-chain invoice data, if available.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TransactionServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetServiceConfig': grpc.unary_unary_rpc_method_handler(
                    servicer.GetServiceConfig,
                    request_deserializer=transaction_dot_v4_dot_transaction__service__pb2.GetServiceConfigRequest.FromString,
                    response_serializer=transaction_dot_v4_dot_transaction__service__pb2.GetServiceConfigResponse.SerializeToString,
            ),
            'GetMinimumKinVersion': grpc.unary_unary_rpc_method_handler(
                    servicer.GetMinimumKinVersion,
                    request_deserializer=transaction_dot_v4_dot_transaction__service__pb2.GetMinimumKinVersionRequest.FromString,
                    response_serializer=transaction_dot_v4_dot_transaction__service__pb2.GetMinimumKinVersionResponse.SerializeToString,
            ),
            'GetRecentBlockhash': grpc.unary_unary_rpc_method_handler(
                    servicer.GetRecentBlockhash,
                    request_deserializer=transaction_dot_v4_dot_transaction__service__pb2.GetRecentBlockhashRequest.FromString,
                    response_serializer=transaction_dot_v4_dot_transaction__service__pb2.GetRecentBlockhashResponse.SerializeToString,
            ),
            'GetMinimumBalanceForRentExemption': grpc.unary_unary_rpc_method_handler(
                    servicer.GetMinimumBalanceForRentExemption,
                    request_deserializer=transaction_dot_v4_dot_transaction__service__pb2.GetMinimumBalanceForRentExemptionRequest.FromString,
                    response_serializer=transaction_dot_v4_dot_transaction__service__pb2.GetMinimumBalanceForRentExemptionResponse.SerializeToString,
            ),
            'GetHistory': grpc.unary_unary_rpc_method_handler(
                    servicer.GetHistory,
                    request_deserializer=transaction_dot_v4_dot_transaction__service__pb2.GetHistoryRequest.FromString,
                    response_serializer=transaction_dot_v4_dot_transaction__service__pb2.GetHistoryResponse.SerializeToString,
            ),
            'SubmitTransaction': grpc.unary_unary_rpc_method_handler(
                    servicer.SubmitTransaction,
                    request_deserializer=transaction_dot_v4_dot_transaction__service__pb2.SubmitTransactionRequest.FromString,
                    response_serializer=transaction_dot_v4_dot_transaction__service__pb2.SubmitTransactionResponse.SerializeToString,
            ),
            'GetTransaction': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTransaction,
                    request_deserializer=transaction_dot_v4_dot_transaction__service__pb2.GetTransactionRequest.FromString,
                    response_serializer=transaction_dot_v4_dot_transaction__service__pb2.GetTransactionResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'kin.agora.transaction.v4.Transaction', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Transaction(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetServiceConfig(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.transaction.v4.Transaction/GetServiceConfig',
            transaction_dot_v4_dot_transaction__service__pb2.GetServiceConfigRequest.SerializeToString,
            transaction_dot_v4_dot_transaction__service__pb2.GetServiceConfigResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetMinimumKinVersion(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.transaction.v4.Transaction/GetMinimumKinVersion',
            transaction_dot_v4_dot_transaction__service__pb2.GetMinimumKinVersionRequest.SerializeToString,
            transaction_dot_v4_dot_transaction__service__pb2.GetMinimumKinVersionResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetRecentBlockhash(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.transaction.v4.Transaction/GetRecentBlockhash',
            transaction_dot_v4_dot_transaction__service__pb2.GetRecentBlockhashRequest.SerializeToString,
            transaction_dot_v4_dot_transaction__service__pb2.GetRecentBlockhashResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetMinimumBalanceForRentExemption(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.transaction.v4.Transaction/GetMinimumBalanceForRentExemption',
            transaction_dot_v4_dot_transaction__service__pb2.GetMinimumBalanceForRentExemptionRequest.SerializeToString,
            transaction_dot_v4_dot_transaction__service__pb2.GetMinimumBalanceForRentExemptionResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetHistory(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.transaction.v4.Transaction/GetHistory',
            transaction_dot_v4_dot_transaction__service__pb2.GetHistoryRequest.SerializeToString,
            transaction_dot_v4_dot_transaction__service__pb2.GetHistoryResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SubmitTransaction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.transaction.v4.Transaction/SubmitTransaction',
            transaction_dot_v4_dot_transaction__service__pb2.SubmitTransactionRequest.SerializeToString,
            transaction_dot_v4_dot_transaction__service__pb2.SubmitTransactionResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetTransaction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.transaction.v4.Transaction/GetTransaction',
            transaction_dot_v4_dot_transaction__service__pb2.GetTransactionRequest.SerializeToString,
            transaction_dot_v4_dot_transaction__service__pb2.GetTransactionResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)
