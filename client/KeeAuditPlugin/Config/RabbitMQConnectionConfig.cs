using KeePass.App.Configuration;

namespace KeeAuditPlugin.Config
{
    public sealed class RabbitMQConnectionConfig
    {
        private readonly AceCustomConfig _config;
        private const string configKey = "KeeAuditPlugin.RabbitMQ";

        public string HostName {
            get => _config.GetString(configKey + ".HostName", "127.0.0.1");
            set => _config.SetString(configKey + ".HostName", value);
        }

        public int Port {
            get => int.Parse(_config.GetString(configKey + ".Port", "5672"));
            set => _config.SetString(configKey + ".Port", value.ToString());
        }

        public string UserName {
            get => _config.GetString(configKey + ".UserName", "guest");
            set => _config.SetString(configKey + ".UserName", value);
        }

        public string Password {
            get => _config.GetString(configKey + ".Password", "guest");
            set => _config.SetString(configKey + ".Password", value);
        }

        public string VirtualHost {
            get => _config.GetString(configKey + ".VirtualHost", "/");
            set => _config.SetString(configKey + ".VirtualHost", value);
        }

        public RabbitMQConnectionConfig(AceCustomConfig config)
        {
            _config = config;
        }
    }
}
