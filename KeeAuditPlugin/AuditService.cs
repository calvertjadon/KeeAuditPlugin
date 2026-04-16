using KeeAuditPlugin.Config;
using RabbitMQ.Client;
using System;

namespace KeeAuditPlugin
{
    public sealed class AuditService
    {
        private Guid _clientId = Guid.NewGuid();
        private IConnection _conn = null;
        private IModel _channel = null;
        private string QueueName => $"keepass.{_clientId}";
        public AuditService(RabbitMQConnectionConfig rabbitMQConfig)
        {

            var connFactory = new ConnectionFactory()
            {
                HostName = rabbitMQConfig.HostName,
                Port = rabbitMQConfig.Port,
                UserName = rabbitMQConfig.UserName,
                Password = rabbitMQConfig.Password,
                VirtualHost = rabbitMQConfig.VirtualHost
            };
            try
            {
                _conn = connFactory.CreateConnection();
            }
            catch (RabbitMQ.Client.Exceptions.BrokerUnreachableException ex)
            {
                throw new FailedToConnectException("Failed to connect to RabbitMQ: " + ex.Message);
            }

            try
            {
                _channel = _conn.CreateModel();
                _channel.QueueDeclare(queue: QueueName,
                                     durable: false,
                                     exclusive: true,
                                     autoDelete: true,
                                     arguments: null);
                _channel.QueueBind(
                    queue: QueueName,
                    exchange: "audit.commands",
                    routingKey: "audit.run.all"
                );
            }
            catch (RabbitMQ.Client.Exceptions.BrokerUnreachableException ex)
            {
                throw new FailedToCreateChannelException("Failed to create channel or declare queue: " + ex.Message);
            }
        }
    }

    public class FailedToConnectException : Exception
    {
        public FailedToConnectException(string message) : base(message) { }
    }

    public class FailedToCreateChannelException : Exception
    {
        public FailedToCreateChannelException(string message) : base(message) { }
    }
}
