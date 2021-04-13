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
        self.SignTransaction = channel.unary_unary(
                '/kin.agora.transaction.v4.Transaction/SignTransaction',
                request_serializer=transaction_dot_v4_dot_transaction__service__pb2.SignTransactionRequest.SerializeToString,
                response_deserializer=transaction_dot_v4_dot_transaction__service__pb2.SignTransactionResponse.FromString,
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

        The subsidizer key returned may vary based on the 'app-index' header.
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

    def SignTransaction(self, request, context):
        """SignTransaction signs the provided transaction, returning the signature to be used.

        The transaction may include the following types of instructions:
        - SplAssociateTokenAccount::CreateAssociatedTokenAccount()
        - SplToken::SetAuthority(CloseAuthority)
        - SplToken::Transfer()
        - SplToken::CloseAccount()
        - Memo::Memo()

        The transaction can be divided into one or more 'regions', which are delineated by
        the memo instruction. Each instruction within a region is considered to be 'related to'
        the memo at the beginning of the region. The first (or only) region may not have a memo.
        For example, if there are instructions before the first memo instruction, or if there
        is no memo at all.

        If an invoice is applied, there must be a memo whose foreign key contains the SHA-226
        of the serialized memo. Additionally, the number of SplToken::Transfer instructions in
        the region _must_ match the number of invoices. Furthermore, the invoice cannot be
        referenced by more than one region.

        Examples:

        Basic Transfer (No Invoice)
        1. SplToken::Transfer()

        Basic Transfer (Invoice)
        1. Memo::Memo(Spend)
        2. SplToken::Transfer()

        Transfer with Cleanup (Sender has token accounts A, B, sending to C)
        1. Memo::Memo(GC) [Optional, 'memoless' region is ok]
        2. SplToken::Transfer(B -> A)
        3. SplToken::CloseAccount(B)
        4. Memo::Memo(Spend)
        5. SplToken::Transfer(A -> C)

        Transfer with Cleanup At End (Sender has token accounts A, B, sending to C)
        1. Memo::Memo(Spend)
        2. SplToken::Transfer(A -> C)
        3. Memo::Memo(GC) [Required, delineate cleanup region from above]
        4. SplToken::Transfer(B -> A)
        5. SplToken::CloseAccount(B)

        Sender Creates Destination (No Invoice)
        1. SplAssociateTokenAccount::CreateAssociatedTokenAccount()
        2. SplToken::SetAuthority(CloseAuthority)
        2. SplToken::Transfer()

        Sender Creates Destination (Invoice)
        1. SplAssociateTokenAccount::CreateAssociatedTokenAccount()
        2. SplToken::SetAuthority(CloseAuthority)
        3. Memo::Memo(Earn)
        4. SplToken::Transfer()
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SubmitTransaction(self, request, context):
        """SubmitTransaction submits a transaction.

        If the transaction is already signed, the SignTransaction webhook will not
        be called.
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
            'SignTransaction': grpc.unary_unary_rpc_method_handler(
                    servicer.SignTransaction,
                    request_deserializer=transaction_dot_v4_dot_transaction__service__pb2.SignTransactionRequest.FromString,
                    response_serializer=transaction_dot_v4_dot_transaction__service__pb2.SignTransactionResponse.SerializeToString,
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
    def SignTransaction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/kin.agora.transaction.v4.Transaction/SignTransaction',
            transaction_dot_v4_dot_transaction__service__pb2.SignTransactionRequest.SerializeToString,
            transaction_dot_v4_dot_transaction__service__pb2.SignTransactionResponse.FromString,
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
