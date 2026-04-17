using KeePassLib;
using KeePassLib.Cryptography;
using System;
using System.Linq.Expressions;

namespace KeeAuditLib.Specifications
{
    public sealed class EntropySpecification : Specification<PwEntry>
    {
        private readonly double _minEntropy;

        public EntropySpecification(double minEntropy)
        {
            _minEntropy = minEntropy;
        }

        private double GetEntropy(PwEntry entry)
        {
            var password = entry.Strings.Get("Password").ReadString();
            return QualityEstimation.EstimatePasswordBits(password.ToCharArray());
        }

        public override Expression<Func<PwEntry, bool>> ToExpression()
        {
            return entry => GetEntropy(entry) >= _minEntropy;
        }
    }
}
