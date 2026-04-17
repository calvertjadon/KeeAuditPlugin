using KeeAuditLib.Specifications;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using System.Collections.Generic;
using System.Linq;

namespace KeeAuditTests
{

    [TestClass]
    public class TestAudit
    {

        private KeePassLib.PwDatabase _database;
        private KeeAuditLib.EntryRepository _entryRepository;
        private KeeAuditLib.DatabaseAuditor _auditor;


        [TestInitialize]
        public void TestInitialize()
        {
            _database = new KeePassLib.PwDatabase();
            _database.RootGroup = new KeePassLib.PwGroup(true, true);
            _entryRepository = new KeeAuditLib.EntryRepository(_database);
            _auditor = new KeeAuditLib.DatabaseAuditor(_entryRepository);
        }

        [TestMethod]
        public void TestMultipleSatisfiedSpecsPasses()
        {
            // arrange
            var mockSpec1 = new MockSpecification(true);
            var mockSpec2 = new MockSpecification(true);

            _auditor.AddSpecification(mockSpec1);
            _auditor.AddSpecification(mockSpec2);

            var entry = new KeePassLib.PwEntry(true, true);
            _database.RootGroup.AddEntry(entry, true);

            // act
            var results = _auditor.Validate().ToList();

            // assert
            var expected = new List<KeeAuditLib.ValidationResult>()
            {
                new KeeAuditLib.ValidationResult(entry, new List<Specification<KeePassLib.PwEntry>>()),
            };
            CollectionAssert.AreEqual(expected, results);
        }

        [TestMethod]
        public void TestMultipleSpecsOneFails()
        {
            // arrange
            var mockSpec1 = new MockSpecification(true);
            var mockSpec2 = new MockSpecification(false);
            _auditor.AddSpecification(mockSpec1);
            _auditor.AddSpecification(mockSpec2);

            var entry = new KeePassLib.PwEntry(true, true);
            _database.RootGroup.AddEntry(entry, true);

            // act
            var results = _auditor.Validate().ToList();

            // assert
            var expected = new List<KeeAuditLib.ValidationResult>()
        {
            new KeeAuditLib.ValidationResult(entry, new List<Specification<KeePassLib.PwEntry>>()
            {
                mockSpec2
            }),
        };
            CollectionAssert.AreEqual(expected, results);
        }
    }
}
