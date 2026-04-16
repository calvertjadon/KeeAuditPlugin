using KeeAuditPlugin.Config;
using KeePass.Plugins;
using KeePassLib.Interfaces;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using System;
using System.Xml.Serialization;

namespace KeeAuditPlugin
{
    public sealed class AuditService
    {
        private readonly Guid _clientId = Guid.NewGuid();
        private readonly IConnection _conn = null;
        private readonly IModel _channel = null;
        private readonly IPluginHost _host = null;
        private readonly IStatusLogger _logger = null;
        private string QueueName => $"keepass.{_clientId}";
        private const string ExchangeName = "audit.commands";

        public AuditService(RabbitMQConnectionConfig rabbitMQConfig, IPluginHost host, IStatusLogger logger)
        {
            _host = host;
            _logger = logger;

            _logger.SetText("Initializing KeeAuditPlugin...", KeePassLib.Interfaces.LogStatusType.Info);

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
                _logger.SetText("Connecting to RabbitMQ...", KeePassLib.Interfaces.LogStatusType.Info);
                _conn = connFactory.CreateConnection();
                _logger.SetText("KeeAuditPlugin initialized successfully.", KeePassLib.Interfaces.LogStatusType.Info);
            }
            catch (RabbitMQ.Client.Exceptions.BrokerUnreachableException ex)
            {
                throw new FailedToConnectException("Failed to connect to RabbitMQ: " + ex.Message);
            }

            _channel = _conn.CreateModel();
            var _ = _channel.QueueDeclare(
                queue: QueueName,
                durable: false,
                exclusive: true,
                autoDelete: true,
                arguments: null
             ) ?? throw new FailedToDeclareQueueException();

            try
            {
                _channel.QueueBind(
                    queue: QueueName,
                    exchange: ExchangeName,
                    routingKey: "audit.run.all"
                );
            }
            catch (RabbitMQ.Client.Exceptions.BrokerUnreachableException ex)
            {
                throw new FailedToBindQueueException(ex.Message);
            }

            var consumer = new EventingBasicConsumer(_channel);
            consumer.Received += (model, ea) =>
            {
                var body = ea.Body.ToArray();
                var message = System.Text.Encoding.UTF8.GetString(body);
                System.Windows.Forms.MessageBox.Show($"Received message: {message}", "KeeAuditPlugin", System.Windows.Forms.MessageBoxButtons.OK, System.Windows.Forms.MessageBoxIcon.Information);
                _channel.BasicAck(deliveryTag: ea.DeliveryTag, multiple: false);
            };

            _channel.BasicConsume(queue: QueueName, autoAck: false, consumer: consumer);
        }


    }


    public class FailedToConnectException : Exception
    {
        public FailedToConnectException(string message) : base(message) { }
    }

    public class FailedToDeclareQueueException : Exception
    {
        public FailedToDeclareQueueException() : base("Failed to declare queue.") { }
    }

    public class  FailedToBindQueueException : Exception 
    {
        public FailedToBindQueueException(string message) : base(message) { }
    }
}
