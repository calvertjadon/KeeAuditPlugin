using System;

namespace KeeAuditTests
{
    internal class MockSpecification : KeeAuditLib.Specifications.Specification<KeePassLib.PwEntry>
    {
        private readonly bool _result;
        public MockSpecification(bool result)
        {
            _result = result;
        }

        public override System.Linq.Expressions.Expression<Func<KeePassLib.PwEntry, bool>> ToExpression()
        {
            return entry => _result;
        }
    }
}
